package pgimpl

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-feature-flag/flag-management/server/dao"
	dbmodel2 "github.com/go-feature-flag/flag-management/server/dao/dbmodel"
	daoerr "github.com/go-feature-flag/flag-management/server/dao/err"
	"github.com/go-feature-flag/flag-management/server/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // we import the driver used by sqlx
)

func NewPostgresDao(connectionString string) (dao.FlagStorage, error) {
	if connectionString == "" {
		return nil, fmt.Errorf("connection string is empty")
	}

	conn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("impossible to connect to the database: %w", err)
	}
	instance := &pgFlagImpl{
		conn: conn,
	}
	return instance, nil
}

type pgFlagImpl struct {
	conn *sqlx.DB
}

// GetFlags return all the flags
func (m *pgFlagImpl) GetFlags(ctx context.Context) ([]model.FeatureFlag, daoerr.DaoError) {
	var f []dbmodel2.FeatureFlag
	err := m.conn.SelectContext(ctx, &f, "SELECT * FROM feature_flags ORDER BY last_updated_date DESC")
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []model.FeatureFlag{}, nil
		}
		return []model.FeatureFlag{}, daoerr.WrapPostgresError(err)
	}
	res := make([]model.FeatureFlag, 0, len(f))
	for _, flag := range f {
		var rules []dbmodel2.Rule
		err := m.conn.SelectContext(ctx, &rules, `SELECT * FROM rules WHERE feature_flag_id = $1 ORDER BY order_index`, flag.ID)
		if err != nil {
			return []model.FeatureFlag{}, daoerr.WrapPostgresError(err)
		}
		convertedFlag, err := flag.ToModelFeatureFlag(rules)
		if err != nil {
			return []model.FeatureFlag{}, daoerr.WrapPostgresError(err)
		}
		res = append(res, convertedFlag)
	}
	return res, nil
}

// GetFlagByID return a flag by its ID
func (m *pgFlagImpl) GetFlagByID(ctx context.Context, id string) (model.FeatureFlag, daoerr.DaoError) {
	var f dbmodel2.FeatureFlag
	err := m.conn.GetContext(ctx, &f, `SELECT * FROM feature_flags WHERE id = $1`, id)
	if err != nil {
		return model.FeatureFlag{}, daoerr.WrapPostgresError(err)
	}

	var rules []dbmodel2.Rule
	errRule := m.conn.SelectContext(
		ctx,
		&rules,
		`SELECT * FROM rules WHERE feature_flag_id = $1 ORDER BY order_index`, f.ID)
	if errRule != nil && !errors.Is(errRule, pgx.ErrNoRows) {
		return model.FeatureFlag{}, daoerr.WrapPostgresError(errRule)
	}

	if convertedFlag, err := f.ToModelFeatureFlag(rules); err != nil {
		return model.FeatureFlag{}, daoerr.WrapPostgresError(err)
	} else {
		return convertedFlag, nil
	}
}

// GetFlagByName return a flag by its name
func (m *pgFlagImpl) GetFlagByName(ctx context.Context, name string) (model.FeatureFlag, daoerr.DaoError) {
	var f dbmodel2.FeatureFlag
	err := m.conn.GetContext(ctx, &f, `SELECT * FROM feature_flags WHERE name = $1`, name)
	if err != nil {
		return model.FeatureFlag{}, daoerr.WrapPostgresError(err)
	}

	var rules []dbmodel2.Rule
	errRule := m.conn.SelectContext(ctx, &rules,
		`SELECT * FROM rules WHERE feature_flag_id = $1 ORDER BY order_index`, f.ID)
	if errRule != nil && !errors.Is(errRule, pgx.ErrNoRows) {
		return model.FeatureFlag{}, daoerr.WrapPostgresError(errRule)
	}
	if convertedFlag, err := f.ToModelFeatureFlag(rules); err != nil {
		return model.FeatureFlag{}, daoerr.WrapPostgresError(err)
	} else {
		return convertedFlag, nil
	}
}

// CreateFlag create a new flag, return the id of the flag
func (m *pgFlagImpl) CreateFlag(ctx context.Context, flag model.FeatureFlag) (string, daoerr.DaoError) {
	dbFeatureFlag, err := dbmodel2.FromModelFeatureFlag(flag)
	if err != nil {
		return "", daoerr.WrapPostgresError(err)
	}

	tx, err := m.conn.Beginx()
	if err != nil {
		return "", daoerr.WrapPostgresError(err)
	}
	defer func() { _ = tx.Commit() }()
	_, err = tx.NamedExecContext(
		ctx,
		`INSERT INTO feature_flags (
                           id,
                           name,
                           description,
                           variations,
                           type,
                           bucketing_key,
                           metadata,
                           track_events,
                           disable,
                           version,
                           created_date,
                           last_updated_date,
                           last_modified_by) 
				VALUES (
				        :id,
				        :name,
				        :description,
				        :variations,
				        :type,
				        :bucketing_key,
				        :metadata,
				        :track_events,
				        :disable,
				        :version,
				        :created_date,
				        :last_updated_date,
				        :last_modified_by)`,
		dbFeatureFlag)
	if err != nil {
		_ = tx.Rollback()
		return "", daoerr.WrapPostgresError(err)
	}

	if flag.DefaultRule == nil {
		return "", daoerr.NewDaoError(daoerr.DefaultRuleRequired, fmt.Errorf("default rule is required"))
	}
	err = m.insertRule(ctx, *flag.DefaultRule, true, dbFeatureFlag.ID, tx, -1)
	if err != nil {
		_ = tx.Rollback()
		return "", daoerr.WrapPostgresError(err)
	}

	if flag.Rules != nil {
		for index, rule := range *flag.Rules {
			err = m.insertRule(ctx, rule, false, dbFeatureFlag.ID, tx, index)
			if err != nil {
				_ = tx.Rollback()
				return "", daoerr.WrapPostgresError(err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return "", daoerr.WrapPostgresError(err)
	}
	return dbFeatureFlag.ID.String(), nil
}

func (m *pgFlagImpl) UpdateFlag(ctx context.Context, flag model.FeatureFlag) daoerr.DaoError {
	dbQuery, errConv := dbmodel2.FromModelFeatureFlag(flag)
	if errConv != nil {
		return daoerr.WrapPostgresError(errConv)
	}
	tx, err := m.conn.Beginx()
	if err != nil {
		return daoerr.WrapPostgresError(err)
	}

	flagOrder := map[string]int{}
	for i, rule := range flag.GetRules() {
		flagOrder[rule.ID] = i
	}

	dbFF, getFlagErr := m.GetFlagByID(ctx, flag.ID)
	if getFlagErr != nil {
		return getFlagErr
	}

	// update default rule
	if flag.DefaultRule == nil {
		return daoerr.NewDaoError(daoerr.DefaultRuleRequired, fmt.Errorf("default rule is required"))
	}

	if err := m.updateRule(ctx, flag.GetDefaultRule(), true, dbQuery.ID, tx, -1); err != nil {
		_ = tx.Rollback
		return daoerr.WrapPostgresError(err)
	}

	listExistingRuleIDs := make(map[string]model.Rule)
	for _, rule := range dbFF.GetRules() {
		listExistingRuleIDs[rule.ID] = rule
	}
	listNewRuleIDs := make(map[string]model.Rule)
	for _, rule := range flag.GetRules() {
		listNewRuleIDs[rule.ID] = rule
	}

	var toDelete, toCreate, toUpdate []string
	for id := range listExistingRuleIDs {
		if _, found := listNewRuleIDs[id]; found {
			toUpdate = append(toUpdate, id)
		} else {
			toDelete = append(toDelete, id)
		}
	}

	for id := range listNewRuleIDs {
		if _, found := listExistingRuleIDs[id]; !found {
			toCreate = append(toCreate, id)
		}
	}

	// Delete rules
	for _, id := range toDelete {
		if _, err := tx.ExecContext(ctx, `DELETE FROM rules WHERE id = $1`, id); err != nil {
			_ = tx.Rollback
			return daoerr.WrapPostgresError(err)
		}
	}

	for _, id := range toCreate {
		rule := listNewRuleIDs[id]
		if err := m.insertRule(ctx, rule, false, dbQuery.ID, tx, flagOrder[id]); err != nil {
			_ = tx.Rollback
			return daoerr.WrapPostgresError(err)
		}
	}

	for _, id := range toUpdate {
		rule := listNewRuleIDs[id]
		if err = m.updateRule(ctx, rule, false, dbQuery.ID, tx, flagOrder[id]); err != nil {
			_ = tx.Rollback
			return daoerr.WrapPostgresError(err)
		}
	}

	// Update the flag
	if _, errTx := tx.NamedExecContext(
		ctx,
		`UPDATE feature_flags SET 
				 name=:name,
				 description=:description,
				 variations=:variations,
				 bucketing_key=:bucketing_key,
				 metadata=:metadata,
				 track_events=:track_events,
				 disable=:disable,
				 version=:version,
				 last_updated_date=:last_updated_date,
				 last_modified_by=:last_modified_by
				WHERE id = :id`,
		dbQuery); err != nil {
		_ = tx.Rollback
		return daoerr.WrapPostgresError(errTx)
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		_ = tx.Rollback
		return daoerr.WrapPostgresError(commitErr)
	}
	return nil
}

func (m *pgFlagImpl) DeleteFlagByID(ctx context.Context, id string) daoerr.DaoError {
	tx, err := m.conn.Beginx()
	if err != nil {
		return daoerr.WrapPostgresError(err)
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM rules WHERE feature_flag_id = $1`, id)
	if err != nil {
		_ = tx.Rollback()
		return daoerr.WrapPostgresError(err)
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM feature_flags WHERE id = $1`, id)
	if err != nil {
		_ = tx.Rollback()
		return daoerr.WrapPostgresError(err)
	}
	commitErr := tx.Commit()
	if commitErr != nil {
		_ = tx.Rollback
		return daoerr.WrapPostgresError(commitErr)
	}
	return nil
}

func (m *pgFlagImpl) insertRule(
	ctx context.Context,
	rule model.Rule,
	isDefault bool,
	featureFlagID uuid.UUID,
	tx *sqlx.Tx,
	orderIndex int) error {
	r, err := dbmodel2.FromModelRule(rule, featureFlagID, isDefault, orderIndex)
	if err != nil {
		return err
	}

	_, errTx := tx.NamedExecContext(
		ctx,
		`INSERT INTO rules (
					id, 
					feature_flag_id, 
					is_default, 
					name,
					query,
					variation_result,
					percentages,
					disable,
					progressive_rollout_initial_variation,
					progressive_rollout_end_variation,
					progressive_rollout_initial_percentage,
					progressive_rollout_end_percentage,
					progressive_rollout_start_date,
					progressive_rollout_end_date,
                   	order_index)
    			VALUES (
					:id,
					:feature_flag_id,
					:is_default,
					:name,
					:query,
					:variation_result,
					:percentages,
					:disable,
					:progressive_rollout_initial_variation,
					:progressive_rollout_end_variation,
					:progressive_rollout_initial_percentage,
					:progressive_rollout_end_percentage,
					:progressive_rollout_start_date,
					:progressive_rollout_end_date,
					:order_index)`,
		r)
	return errTx
}

func (m *pgFlagImpl) updateRule(
	ctx context.Context,
	rule model.Rule,
	isDefault bool,
	featureFlagID uuid.UUID,
	tx *sqlx.Tx, orderIndex int) error {
	r, err := dbmodel2.FromModelRule(rule, featureFlagID, isDefault, orderIndex)
	if err != nil {
		return err
	}

	_, errTx := tx.NamedExecContext(ctx,
		`UPDATE rules SET 
                 name=:name, 
                 query=:query,
                 variation_result=:variation_result,
                 percentages=:percentages,
                 disable=:disable,
                 progressive_rollout_initial_variation=:progressive_rollout_initial_variation,
                 progressive_rollout_end_variation=:progressive_rollout_end_variation,
                 progressive_rollout_initial_percentage=:progressive_rollout_initial_percentage,
                 progressive_rollout_end_percentage=:progressive_rollout_end_percentage,
                 progressive_rollout_start_date=:progressive_rollout_start_date,
                 progressive_rollout_end_date=:progressive_rollout_end_date
             WHERE id=:id`, r)

	return errTx
}

func (m *pgFlagImpl) Ping() daoerr.DaoError {
	if m.conn == nil {
		return daoerr.NewDaoError(daoerr.DatabaseNotInitialized, errors.New("database connection is nil"))
	}
	err := m.conn.Ping()
	if err != nil {
		return daoerr.WrapPostgresError(err)
	}
	return nil
}
