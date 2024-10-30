package pgimpl_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-feature-flag/app-api/dao"
	daoerr "github.com/go-feature-flag/app-api/dao/err"
	"github.com/go-feature-flag/app-api/dao/pgimpl"
	"github.com/go-feature-flag/app-api/model"
	"github.com/go-feature-flag/app-api/testutils"
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

func getPostgresDao(t *testing.T, pgContainer *testcontainerPostgres.PostgresContainer) dao.Flags {
	mappedPort, err := pgContainer.MappedPort(context.Background(), "5432")
	require.NoError(t, err, "Failed to get mapped port")
	port, _ := strconv.Atoi(mappedPort.Port())
	d, err := pgimpl.NewPostgresDao("localhost", port, databaseName, databaseUsername, databasePassword)
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
			wantErrMessage: "host is empty",
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
			wantErrMessage: "invalid port: port is 0",
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
			wantErrMessage: "dbName is empty",
		},
		{
			name: "should return an error for an empty user",
			args: args{
				host:     "localhost",
				port:     5432,
				dbName:   "test",
				user:     "",
				password: "password",
			},
			wantErr:        assert.Error,
			wantErrMessage: "username is empty",
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
			wantErrMessage: "password is empty",
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
			wantErrMessage: "impossible to connect to the database: dial tcp [::1]:5432: connect: connection refused",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := pgimpl.NewPostgresDao(tt.args.host, tt.args.port, tt.args.dbName, tt.args.user, tt.args.password)
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

			_, err := pgDao.GetFlagByID(context.TODO(), tt.id)
			tt.wantErr(t, err)
			assert.Equal(t, tt.wantDaoErr, err)
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

			_, err := pgDao.GetFlagByName(context.TODO(), tt.flagName)
			tt.wantErr(t, err)
			assert.Equal(t, tt.wantDaoErr, err)
		})
	}

}
