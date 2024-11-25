//go:build docker
// +build docker

package pgimpl_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-feature-flag/flag-management/server/dao"
	"github.com/go-feature-flag/flag-management/server/dao/pgimpl"
	"github.com/go-feature-flag/flag-management/server/testutils"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-feature-flag/flag-management/server/model"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // we import the driver used by sqlx
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	testcontainerPostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	databaseName     = "test"
	databaseUsername = "user"
	databasePassword = "password"
)

// setupTest will start a Postgres container and run the migrations
func setupTest(t *testing.T, sqlFileToInsert []string) (*testcontainerPostgres.PostgresContainer, *sqlx.DB) {
	pgContainer, err := testcontainerPostgres.Run(context.Background(),
		"postgres:16-alpine",
		testcontainerPostgres.WithDatabase(databaseName),
		testcontainerPostgres.WithUsername(databaseUsername),
		testcontainerPostgres.WithPassword(databasePassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	require.NoError(t, err, "Failed to start Postgres container")
	connectionString, err := pgContainer.ConnectionString(context.Background())
	require.NoError(t, err, "Failed to get connection string")
	connectionString = fmt.Sprintf("%s%s", connectionString, "sslmode=disable")
	conn, err := sqlx.Connect("postgres", connectionString)
	require.NoError(t, err, "Failed to connect to Postgres")
	driver, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	require.NoError(t, err, "Failed to create Postgres driver instance")
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../database_migration",
		"postgres", driver)
	err = m.Up()
	require.NoError(t, err)
	if sqlFileToInsert != nil && len(sqlFileToInsert) > 0 {
		for _, initDataPath := range sqlFileToInsert {
			content, err := os.ReadFile(initDataPath)
			require.NoError(t, err, "Failed to read initial data file")
			_, err = conn.Exec(string(content))
			require.NoErrorf(t, err, "error while inserting data from file %s", initDataPath)
		}
	}
	return pgContainer, conn
}

// tearDownTest will stop the Postgres container
func tearDownTest(t *testing.T, pgContainer *testcontainerPostgres.PostgresContainer, conn *sqlx.DB) {
	driver, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../database_migration",
		"postgres", driver)
	err = m.Down()
	require.NoError(t, err)

	closeTime := 10 * time.Second
	err = pgContainer.Stop(context.Background(), &closeTime)
	require.NoError(t, err, "Failed to stop Postgres container")
}

func getPostgresDao(t *testing.T, pgContainer *testcontainerPostgres.PostgresContainer) dao.FlagStorage {
	mappedPort, err := pgContainer.MappedPort(context.Background(), "5432")
	require.NoError(t, err, "Failed to get mapped port")
	port, _ := strconv.Atoi(mappedPort.Port())
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", databaseUsername, databasePassword, "localhost", port, databaseName)
	d, err := pgimpl.NewPostgresDao(connectionString)
	require.NoError(t, err, "Failed to create Postgres DAO")
	return d
}

func TestNewPostgresDao(t *testing.T) {
	type args struct {
		host     string
		port     int
		dbName   string
		user     string
		password string
	}
	tests := []struct {
		name           string
		args           args
		wantErr        assert.ErrorAssertionFunc
		wantErrMessage string
	}{
		{
			name: "should return an error for an empty host",
			args: args{
				host:     "",
				port:     5432,
				dbName:   "test",
				user:     "user",
				password: "password",
			},
			wantErr:        assert.Error,
			wantErrMessage: "impossible to connect to the database: pq: password authentication failed for user \"user\"",
		},
		{
			name: "should return an error for an empty port",
			args: args{
				host:     "localhost",
				dbName:   "test",
				user:     "user",
				password: "password",
			},
			wantErr:        assert.Error,
			wantErrMessage: "impossible to connect to the database: dial tcp [::1]:0: connect: can't assign requested address",
		},
		{
			name: "should return an error for an empty dbName",
			args: args{
				host:     "localhost",
				port:     5432,
				dbName:   "",
				user:     "user",
				password: "password",
			},
			wantErr:        assert.Error,
			wantErrMessage: "impossible to connect to the database: pq: password authentication failed for user \"user\"",
		},
		{
			name: "should return an error for an empty password",
			args: args{
				host:     "localhost",
				port:     5432,
				dbName:   "test",
				user:     "user",
				password: "",
			},
			wantErr:        assert.Error,
			wantErrMessage: "impossible to connect to the database: pq: password authentication failed for user \"user\"",
		},
		{
			name: "should return an error for a not available DB",
			args: args{
				host:     "localhost",
				port:     5432,
				dbName:   "test",
				user:     "user",
				password: "password",
			},
			wantErr:        assert.Error,
			wantErrMessage: "impossible to connect to the database: pq: password authentication failed for user \"user\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", tt.args.user, tt.args.password, tt.args.host, tt.args.port, tt.args.dbName)
			_, err := pgimpl.NewPostgresDao(connectionString)
			tt.wantErr(t, err)
			if err != nil {
				assert.Equal(t, tt.wantErrMessage, err.Error())
			}
		})
	}
}

func TestGetFlags(t *testing.T) {
	tests := []struct {
		name      string
		initFiles []string
		wantErr   assert.ErrorAssertionFunc
		want      []model.FeatureFlag
	}{
		{
			name:      "should return not error and return 1 flag",
			initFiles: []string{"./testdata/initial_data.sql"},
			wantErr:   assert.NoError,
			want: []model.FeatureFlag{
				{
					ID:          "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
					Name:        "my-feature-flag",
					Description: testutils.String("This is a feature flag"),
					Variations: &map[string]interface{}{
						"variationA": "valueA",
						"variationB": "valueB",
					},
					VariationType: "string",
					BucketingKey:  nil,
					Metadata: &map[string]interface{}{
						"key": "value",
					},
					TrackEvents:     testutils.Bool(true),
					Disable:         testutils.Bool(false),
					Version:         testutils.String("1.0.0"),
					CreatedDate:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
					LastUpdatedDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
					LastModifiedBy:  "admin",
					DefaultRule: &model.Rule{
						ID:              "1cb941f2-adb4-460f-9259-b4416c90e9e1",
						Name:            "default-rule",
						VariationResult: testutils.String("variationA"),
					},
					Rules: &[]model.Rule{
						{
							ID:      "9f82fe80-b4b6-426a-869a-e4436de66d0d",
							Name:    "rule 2",
							Query:   "targetingKey eq \"1234\"",
							Disable: true,
							ProgressiveRollout: &model.ProgressiveRollout{
								Initial: &model.ProgressiveRolloutStep{
									Variation:  testutils.String("variationA"),
									Percentage: testutils.Float64(0),
									Date:       testutils.Time(time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0))),
								},
								End: &model.ProgressiveRolloutStep{
									Variation:  testutils.String("variationB"),
									Percentage: testutils.Float64(100),
									Date:       testutils.Time(time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0))),
								},
							},
						},
						{
							ID:      "546939a9-6df8-4a0b-b9cf-1d69ff300eb5",
							Name:    "rule 1",
							Query:   "targetingKey eq \"valueA\"",
							Disable: false,
							Percentages: &map[string]float64{
								"variationA": 10,
								"variationB": 90,
							},
						},
					},
				},
			},
		},
		{
			name:      "should return an error if the flags table is empty",
			initFiles: []string{},
			wantErr:   assert.NoError,
			want:      []model.FeatureFlag{},
		},
		{
			name:      "should return an error if no default rule is present",
			initFiles: []string{"./testdata/initial_data_without_default_rule.sql"},
			wantErr:   assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pgContainer, conn := setupTest(t, tt.initFiles)
			defer tearDownTest(t, pgContainer, conn)
			pgDao := getPostgresDao(t, pgContainer)

			flags, err := pgDao.GetFlags(context.TODO())
			tt.wantErr(t, err)
			if err == nil {
				assert.Equal(t, tt.want, flags)
			}
		})
	}

}

func TestGetFlagByID(t *testing.T) {
	tests := []struct {
		name       string
		initFiles  []string
		id         string
		wantErr    assert.ErrorAssertionFunc
		wantDaoErr daoerr.DaoError
		want       model.FeatureFlag
	}{
		{
			name:       "should return an error if the flag is not found",
			initFiles:  []string{"./testdata/initial_data.sql"},
			wantErr:    assert.Error,
			id:         "546939a9-6df8-4a0b-b9cf-1d69ff300eb5",
			wantDaoErr: daoerr.NewDaoError(daoerr.NotFound, errors.New("sql: no rows in result set")),
		}, {
			name:      "should not return an error if the flag is found",
			initFiles: []string{"./testdata/initial_data.sql"},
			wantErr:   assert.NoError,
			id:        "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
			want: model.FeatureFlag{
				ID:          "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
				Name:        "my-feature-flag",
				Description: testutils.String("This is a feature flag"),
				Variations: &map[string]interface{}{
					"variationA": "valueA",
					"variationB": "valueB",
				},
				VariationType: "string",
				BucketingKey:  nil,
				Metadata: &map[string]interface{}{
					"key": "value",
				},
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("1.0.0"),
				CreatedDate:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastUpdatedDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastModifiedBy:  "admin",
				DefaultRule: &model.Rule{
					ID:              "1cb941f2-adb4-460f-9259-b4416c90e9e1",
					Name:            "default-rule",
					VariationResult: testutils.String("variationA"),
				},
				Rules: &[]model.Rule{
					{
						ID:      "9f82fe80-b4b6-426a-869a-e4436de66d0d",
						Name:    "rule 2",
						Query:   "targetingKey eq \"1234\"",
						Disable: true,
						ProgressiveRollout: &model.ProgressiveRollout{
							Initial: &model.ProgressiveRolloutStep{
								Variation:  testutils.String("variationA"),
								Percentage: testutils.Float64(0),
								Date:       testutils.Time(time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0))),
							},
							End: &model.ProgressiveRolloutStep{
								Variation:  testutils.String("variationB"),
								Percentage: testutils.Float64(100),
								Date:       testutils.Time(time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0))),
							},
						},
					},
					{
						ID:      "546939a9-6df8-4a0b-b9cf-1d69ff300eb5",
						Name:    "rule 1",
						Query:   "targetingKey eq \"valueA\"",
						Disable: false,
						Percentages: &map[string]float64{
							"variationA": 10,
							"variationB": 90,
						},
					},
				},
			},
		}, {
			name:       "should return an error for an invalid flag",
			initFiles:  []string{"./testdata/initial_data_without_default_rule.sql"},
			wantErr:    assert.Error,
			id:         "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
			wantDaoErr: daoerr.NewDaoError(daoerr.UnknownError, errors.New("default rule is required")),
		}, {
			name:       "should return an error if no rule is found for the flag",
			initFiles:  []string{"./testdata/initial_data_without_rules.sql"},
			wantErr:    assert.Error,
			id:         "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
			wantDaoErr: daoerr.NewDaoError(daoerr.UnknownError, errors.New("default rule is required")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pgContainer, conn := setupTest(t, tt.initFiles)
			defer tearDownTest(t, pgContainer, conn)
			pgDao := getPostgresDao(t, pgContainer)

			got, err := pgDao.GetFlagByID(context.TODO(), tt.id)
			tt.wantErr(t, err)
			if err != nil {
				assert.Equal(t, tt.wantDaoErr, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}

}

func TestGetFlagByName(t *testing.T) {
	tests := []struct {
		name       string
		initFiles  []string
		flagName   string
		wantErr    assert.ErrorAssertionFunc
		wantDaoErr daoerr.DaoError
		want       model.FeatureFlag
	}{
		{
			name:       "should return an error if the flag is not found",
			initFiles:  []string{"./testdata/initial_data.sql"},
			wantErr:    assert.Error,
			flagName:   "unknown-flag",
			wantDaoErr: daoerr.NewDaoError(daoerr.NotFound, errors.New("sql: no rows in result set")),
		}, {
			name:      "should not return an error if the flag is found",
			initFiles: []string{"./testdata/initial_data.sql"},
			wantErr:   assert.NoError,
			flagName:  "my-feature-flag",
			want: model.FeatureFlag{
				ID:          "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
				Name:        "my-feature-flag",
				Description: testutils.String("This is a feature flag"),
				Variations: &map[string]interface{}{
					"variationA": "valueA",
					"variationB": "valueB",
				},
				VariationType: "string",
				BucketingKey:  nil,
				Metadata: &map[string]interface{}{
					"key": "value",
				},
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("1.0.0"),
				CreatedDate:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastUpdatedDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastModifiedBy:  "admin",
				DefaultRule: &model.Rule{
					ID:              "1cb941f2-adb4-460f-9259-b4416c90e9e1",
					Name:            "default-rule",
					VariationResult: testutils.String("variationA"),
				},
				Rules: &[]model.Rule{
					{
						ID:      "9f82fe80-b4b6-426a-869a-e4436de66d0d",
						Name:    "rule 2",
						Query:   "targetingKey eq \"1234\"",
						Disable: true,
						ProgressiveRollout: &model.ProgressiveRollout{
							Initial: &model.ProgressiveRolloutStep{
								Variation:  testutils.String("variationA"),
								Percentage: testutils.Float64(0),
								Date:       testutils.Time(time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0))),
							},
							End: &model.ProgressiveRolloutStep{
								Variation:  testutils.String("variationB"),
								Percentage: testutils.Float64(100),
								Date:       testutils.Time(time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0))),
							},
						},
					},
					{
						ID:      "546939a9-6df8-4a0b-b9cf-1d69ff300eb5",
						Name:    "rule 1",
						Query:   "targetingKey eq \"valueA\"",
						Disable: false,
						Percentages: &map[string]float64{
							"variationA": 10,
							"variationB": 90,
						},
					},
				},
			},
		}, {
			name:       "should return an error for an invalid flag",
			initFiles:  []string{"./testdata/initial_data_without_default_rule.sql"},
			wantErr:    assert.Error,
			flagName:   "my-feature-flag",
			wantDaoErr: daoerr.NewDaoError(daoerr.UnknownError, errors.New("default rule is required")),
		}, {
			name:       "should return an error if no rule found for the flag",
			initFiles:  []string{"./testdata/initial_data_without_rules.sql"},
			wantErr:    assert.Error,
			flagName:   "my-feature-flag",
			wantDaoErr: daoerr.NewDaoError(daoerr.UnknownError, errors.New("default rule is required")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pgContainer, conn := setupTest(t, tt.initFiles)
			defer tearDownTest(t, pgContainer, conn)
			pgDao := getPostgresDao(t, pgContainer)

			got, err := pgDao.GetFlagByName(context.TODO(), tt.flagName)
			if err != nil {
				tt.wantErr(t, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
			assert.Equal(t, tt.wantDaoErr, err)
		})
	}

}

func TestCreateFlag(t *testing.T) {
	tests := []struct {
		name         string
		flagToCreate model.FeatureFlag
		initFiles    []string
		wantErr      assert.ErrorAssertionFunc
		wantDaoErr   daoerr.DaoError
		wantUUID     string
	}{
		{
			name:      "should create a new flag",
			initFiles: []string{"./testdata/initial_data.sql"},
			wantErr:   assert.NoError,
			wantUUID:  "6e0133ab-c262-4a0e-9eb1-79173c214921",
			flagToCreate: model.FeatureFlag{
				ID:          "6e0133ab-c262-4a0e-9eb1-79173c214921",
				Name:        "my-new-feature-flag",
				Description: testutils.String("This is a feature flag"),
				Variations: &map[string]interface{}{
					"variationA": 10,
					"variationB": 120,
				},
				VariationType: "integer",
				BucketingKey:  nil,
				Metadata: &map[string]interface{}{
					"key": "value",
				},
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("0.0.1"),
				CreatedDate:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastUpdatedDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastModifiedBy:  "foo",
				DefaultRule: &model.Rule{
					ID:              "6761c19f-1b74-49f1-9101-4c4aaa7e89e2",
					Name:            "default-rule",
					VariationResult: testutils.String("variationA"),
				},
				Rules: &[]model.Rule{
					{
						ID:      "2300e15f-7755-4e39-b7c3-957f88f968bd",
						Name:    "rule 2",
						Query:   "targetingKey eq \"1234\"",
						Disable: true,
						ProgressiveRollout: &model.ProgressiveRollout{
							Initial: &model.ProgressiveRolloutStep{
								Variation:  testutils.String("variationA"),
								Percentage: testutils.Float64(0),
								Date:       testutils.Time(time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0))),
							},
							End: &model.ProgressiveRolloutStep{
								Variation:  testutils.String("variationB"),
								Percentage: testutils.Float64(100),
								Date:       testutils.Time(time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0))),
							},
						},
					},
					{
						ID:      "8db65089-6785-4965-84a9-5d055469854b",
						Name:    "rule 1",
						Query:   "targetingKey eq \"valueA\"",
						Disable: false,
						Percentages: &map[string]float64{
							"variationA": 10,
							"variationB": 90,
						},
					},
				},
			},
		}, {
			name:       "should return an error if UUID not valid",
			initFiles:  []string{"./testdata/initial_data.sql"},
			wantErr:    assert.Error,
			wantDaoErr: daoerr.NewDaoError(daoerr.InvalidUUID, errors.New("invalid UUID length: 12")),
			flagToCreate: model.FeatureFlag{
				ID:          "invalid-uuid",
				Name:        "my-new-feature-flag",
				Description: testutils.String("This is a feature flag"),
				Variations: &map[string]interface{}{
					"variationA": 10,
					"variationB": 120,
				},
				VariationType: "integer",
				BucketingKey:  nil,
				Metadata: &map[string]interface{}{
					"key": "value",
				},
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("0.0.1"),
				CreatedDate:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastUpdatedDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastModifiedBy:  "foo",
				DefaultRule: &model.Rule{
					ID:              "6761c19f-1b74-49f1-9101-4c4aaa7e89e2",
					Name:            "default-rule",
					VariationResult: testutils.String("variationA"),
				},
			},
		}, {
			name:       "should return an error if no default rule is present",
			initFiles:  []string{"./testdata/initial_data.sql"},
			wantErr:    assert.Error,
			wantDaoErr: daoerr.NewDaoError(daoerr.DefaultRuleRequired, errors.New("default rule is required")),
			flagToCreate: model.FeatureFlag{
				ID:          "6e0133ab-c262-4a0e-9eb1-79173c214921",
				Name:        "my-new-feature-flag",
				Description: testutils.String("This is a feature flag"),
				Variations: &map[string]interface{}{
					"variationA": 10,
					"variationB": 120,
				},
				VariationType: "integer",
				BucketingKey:  nil,
				Metadata: &map[string]interface{}{
					"key": "value",
				},
				TrackEvents:     testutils.Bool(true),
				Disable:         testutils.Bool(false),
				Version:         testutils.String("0.0.1"),
				CreatedDate:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastUpdatedDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastModifiedBy:  "foo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pgContainer, conn := setupTest(t, tt.initFiles)
			defer tearDownTest(t, pgContainer, conn)
			pgDao := getPostgresDao(t, pgContainer)

			got, err := pgDao.CreateFlag(context.TODO(), tt.flagToCreate)
			tt.wantErr(t, err)
			if err == nil {
				assert.Equal(t, tt.wantUUID, got)
			} else {
				assert.Equal(t, tt.wantDaoErr.Code(), err.Code())
				assert.Equal(t, tt.wantDaoErr.Error(), err.Error())
			}
		})
	}

}

func TestDeleteFlagByID(t *testing.T) {
	tests := []struct {
		name       string
		initFiles  []string
		id         string
		wantErr    assert.ErrorAssertionFunc
		wantDaoErr daoerr.DaoError
		want       model.FeatureFlag
	}{
		{
			name:      "should return not error if the flag is not found",
			initFiles: []string{"./testdata/initial_data.sql"},
			wantErr:   assert.NoError,
			id:        "546939a9-6df8-4a0b-b9cf-1d69ff300eb5",
		}, {
			name:      "should delete the flag if the flag is found",
			initFiles: []string{"./testdata/initial_data.sql"},
			wantErr:   assert.NoError,
			id:        "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pgContainer, conn := setupTest(t, tt.initFiles)
			defer tearDownTest(t, pgContainer, conn)
			pgDao := getPostgresDao(t, pgContainer)

			err := pgDao.DeleteFlagByID(context.TODO(), tt.id)
			tt.wantErr(t, err)
			if err != nil {
				assert.Equal(t, tt.wantDaoErr, err)
			} else {
				_, err := pgDao.GetFlagByID(context.TODO(), tt.id)
				assert.Equal(t, err.Code(), daoerr.NotFound)
			}
		})
	}

}

func TestUpdateFlag(t *testing.T) {
	tests := []struct {
		name         string
		flagToUpdate model.FeatureFlag
		initFiles    []string
		want         model.FeatureFlag
		wantErr      assert.ErrorAssertionFunc
		wantDaoErr   daoerr.DaoError
	}{
		{
			name:      "should update an existing flag",
			initFiles: []string{"./testdata/initial_data.sql"},
			wantErr:   assert.NoError,
			flagToUpdate: model.FeatureFlag{
				ID:          "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
				Name:        "updated-name",
				Description: testutils.String("Updated description"),
				Variations: &map[string]interface{}{
					"variationC": "110",
					"variationD": "10",
				},
				VariationType: "string",
				BucketingKey:  testutils.String("teamID"),
				Metadata: &map[string]interface{}{
					"metadata1": "1",
					"metadata2": "2",
					"metadata3": "3",
				},
				TrackEvents:     testutils.Bool(false),
				Disable:         testutils.Bool(true),
				Version:         testutils.String("0.0.1"),
				CreatedDate:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastUpdatedDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastModifiedBy:  "bar",
				DefaultRule: &model.Rule{
					ID:   "1cb941f2-adb4-460f-9259-b4416c90e9e1",
					Name: "default-rule",
					Percentages: &map[string]float64{
						"variationC": 10,
						"variationD": 90,
					},
				},
				Rules: &[]model.Rule{
					{
						ID:      "8db65089-6785-4965-84a9-5d055469854b",
						Name:    "rule FOO",
						Query:   "targetingKey eq \"91011\"",
						Disable: false,
						Percentages: &map[string]float64{
							"variationA": 10,
							"variationB": 90,
						},
					},
					{
						ID:              "2300e15f-7755-4e39-b7c3-957f88f968bd",
						Name:            "rule BAR",
						Query:           "targetingKey eq \"5678\"",
						Disable:         true,
						VariationResult: testutils.String("variationD"),
					},
				},
			},
			want: model.FeatureFlag{
				ID:          "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
				Name:        "updated-name",
				Description: testutils.String("Updated description"),
				Variations: &map[string]interface{}{
					"variationC": "110",
					"variationD": "10",
				},
				VariationType: "string",
				BucketingKey:  testutils.String("teamID"),
				Metadata: &map[string]interface{}{
					"metadata1": "1",
					"metadata2": "2",
					"metadata3": "3",
				},
				TrackEvents:     testutils.Bool(false),
				Disable:         testutils.Bool(true),
				Version:         testutils.String("0.0.1"),
				CreatedDate:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastUpdatedDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastModifiedBy:  "bar",
				DefaultRule: &model.Rule{
					ID:   "1cb941f2-adb4-460f-9259-b4416c90e9e1",
					Name: "default-rule",
					Percentages: &map[string]float64{
						"variationC": 10,
						"variationD": 90,
					},
				},
				Rules: &[]model.Rule{
					{
						ID:      "8db65089-6785-4965-84a9-5d055469854b",
						Name:    "rule FOO",
						Query:   "targetingKey eq \"91011\"",
						Disable: false,
						Percentages: &map[string]float64{
							"variationA": 10,
							"variationB": 90,
						},
					},
					{
						ID:              "2300e15f-7755-4e39-b7c3-957f88f968bd",
						Name:            "rule BAR",
						Query:           "targetingKey eq \"5678\"",
						Disable:         true,
						VariationResult: testutils.String("variationD"),
					},
				},
			},
		}, {
			name:       "should return an error if UUID not valid",
			initFiles:  []string{"./testdata/initial_data.sql"},
			wantErr:    assert.Error,
			wantDaoErr: daoerr.NewDaoError(daoerr.InvalidUUID, errors.New("invalid UUID length: 12")),
			flagToUpdate: model.FeatureFlag{
				ID: "invalid-uuid",
			},
		}, {
			name:       "should return an error if no default rule is present",
			initFiles:  []string{"./testdata/initial_data.sql"},
			wantErr:    assert.Error,
			wantDaoErr: daoerr.NewDaoError(daoerr.DefaultRuleRequired, errors.New("default rule is required")),
			flagToUpdate: model.FeatureFlag{
				ID:          "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
				Name:        "my-new-feature-flag",
				Description: testutils.String("This is a feature flag"),
				Variations: &map[string]interface{}{
					"variationA": 10,
					"variationB": 120,
				},
				VariationType: "integer",
				BucketingKey:  nil,
				Metadata: &map[string]interface{}{
					"key": "value",
				},
				LastUpdatedDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastModifiedBy:  "foo",
			},
		}, {
			name:      "should be able to add reorder and remove rules",
			initFiles: []string{"./testdata/initial_data.sql"},
			wantErr:   assert.NoError,
			flagToUpdate: model.FeatureFlag{
				ID:          "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
				Name:        "updated-name",
				Description: testutils.String("Updated description"),
				Variations: &map[string]interface{}{
					"variationC": "110",
					"variationD": "10",
				},
				VariationType: "string",
				BucketingKey:  testutils.String("teamID"),
				Metadata: &map[string]interface{}{
					"metadata1": "1",
					"metadata2": "2",
					"metadata3": "3",
				},
				TrackEvents:     testutils.Bool(false),
				Disable:         testutils.Bool(true),
				Version:         testutils.String("0.0.1"),
				CreatedDate:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastUpdatedDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastModifiedBy:  "bar",
				DefaultRule: &model.Rule{
					ID:   "1cb941f2-adb4-460f-9259-b4416c90e9e1",
					Name: "default-rule",
					Percentages: &map[string]float64{
						"variationC": 10,
						"variationD": 90,
					},
				},
				Rules: &[]model.Rule{
					{
						ID:              "546939a9-6df8-4a0b-b9cf-1d69ff300eb5",
						Name:            "rule FOO",
						Query:           "targetingKey eq \"91011\"",
						Disable:         false,
						VariationResult: testutils.String("variationC"),
					},
					{
						ID:              "0734784b-d44e-43f7-a0b3-fa85e1c240df",
						Name:            "rule RANDOM 1",
						Query:           "targetingKey eq \"toto\"",
						Disable:         false,
						VariationResult: testutils.String("variationD"),
					},
				},
			},
			want: model.FeatureFlag{
				ID:          "69aa10ec-ec3e-4139-8cdf-6902a5746e2d",
				Name:        "updated-name",
				Description: testutils.String("Updated description"),
				Variations: &map[string]interface{}{
					"variationC": "110",
					"variationD": "10",
				},
				VariationType: "string",
				BucketingKey:  testutils.String("teamID"),
				Metadata: &map[string]interface{}{
					"metadata1": "1",
					"metadata2": "2",
					"metadata3": "3",
				},
				TrackEvents:     testutils.Bool(false),
				Disable:         testutils.Bool(true),
				Version:         testutils.String("0.0.1"),
				CreatedDate:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastUpdatedDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)),
				LastModifiedBy:  "bar",
				DefaultRule: &model.Rule{
					ID:   "1cb941f2-adb4-460f-9259-b4416c90e9e1",
					Name: "default-rule",
					Percentages: &map[string]float64{
						"variationC": 10,
						"variationD": 90,
					},
				},
				Rules: &[]model.Rule{
					{
						ID:              "546939a9-6df8-4a0b-b9cf-1d69ff300eb5",
						Name:            "rule FOO",
						Query:           "targetingKey eq \"91011\"",
						Disable:         false,
						VariationResult: testutils.String("variationC"),
					},
					{
						ID:              "0734784b-d44e-43f7-a0b3-fa85e1c240df",
						Name:            "rule RANDOM 1",
						Query:           "targetingKey eq \"toto\"",
						Disable:         false,
						VariationResult: testutils.String("variationD"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pgContainer, conn := setupTest(t, tt.initFiles)
			defer tearDownTest(t, pgContainer, conn)
			pgDao := getPostgresDao(t, pgContainer)

			err := pgDao.UpdateFlag(context.TODO(), tt.flagToUpdate)
			tt.wantErr(t, err)
			if err == nil {
				got, err := pgDao.GetFlagByID(context.TODO(), tt.flagToUpdate.ID)
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.Equal(t, tt.wantDaoErr.Error(), err.Error())
				assert.Equal(t, tt.wantDaoErr.Code(), err.Code())
			}
		})
	}
}

func TestPingSuccess(t *testing.T) {
	pgContainer, conn := setupTest(t, []string{})
	defer tearDownTest(t, pgContainer, conn)
	pgDao := getPostgresDao(t, pgContainer)
	assert.NoError(t, pgDao.Ping())
}
