package daoerr_test

import (
	"database/sql"
	"errors"
	"testing"

	daoerr "github.com/go-feature-flag/app-api/dao/err"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestWrapPostgresError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want daoerr.DaoError
	}{
		{
			name: "should return not found error",
			err:  sql.ErrNoRows,
			want: daoerr.NewDaoError(daoerr.NotFound, sql.ErrNoRows),
		},
		{
			name: "should return invalid uuid error",
			err:  &pq.Error{Code: "22P02"},
			want: daoerr.NewDaoError(daoerr.InvalidUUID, &pq.Error{Code: "22P02"}),
		},
		{
			name: "should return an unknown error",
			err:  errors.New("random error"),
			want: daoerr.NewDaoError(daoerr.UnknownError, errors.New("random error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, daoerr.WrapPostgresError(tt.err))
		})
	}
}
