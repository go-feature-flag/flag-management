package daoErr

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

// WrapPostgresError wraps a postgres error into a DaoError to have a DB agnostic error handling in the handlers
func WrapPostgresError(err error) DaoError {
	if errors.Is(err, sql.ErrNoRows) {
		return NewDaoError(NotFound, err)
	}
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "22P02" {
		return NewDaoError(InvalidUUID, err)
	}
	return NewDaoError(UnknownError, err)
}
