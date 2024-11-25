package daoerr_test

import (
	"database/sql"
	"errors"
	daoerr2 "github.com/go-feature-flag/flag-management/server/dao/err"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestWrapPostgresError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want daoerr2.DaoError
	}{
		{
			name: "should return not found error",
			err:  sql.ErrNoRows,
			want: daoerr2.NewDaoError(daoerr2.NotFound, sql.ErrNoRows),
		},
		{
			name: "should return invalid uuid error",
			err:  &pq.Error{Code: "22P02"},
			want: daoerr2.NewDaoError(daoerr2.InvalidUUID, &pq.Error{Code: "22P02"}),
		},
		{
			name: "should return an unknown error",
			err:  errors.New("random error"),
			want: daoerr2.NewDaoError(daoerr2.UnknownError, errors.New("random error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, daoerr2.WrapPostgresError(tt.err))
		})
	}
}
