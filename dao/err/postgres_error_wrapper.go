package daoerr

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// WrapPostgresError wraps a postgres error into a DaoError to have a DB agnostic error handling in the handlers
func WrapPostgresError(err error) DaoError {
	var pqErr *pq.Error
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewDaoError(NotFound, err)
	case uuid.IsInvalidLengthError(err), errors.As(err, &pqErr) && pqErr.Code == "22P02":
		return NewDaoError(InvalidUUID, err)
	default:
		return NewDaoError(UnknownError, err)

	}
}
