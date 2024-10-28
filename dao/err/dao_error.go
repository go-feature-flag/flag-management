package daoerr

type DaoErrorCode string

const (
	NotFound               DaoErrorCode = "NOT_FOUND"
	InvalidUUID            DaoErrorCode = "INVALID_UUID"
	ConversionError        DaoErrorCode = "CONVERSION_ERROR"
	DefaultRuleRequired    DaoErrorCode = "DEFAULT_RULE_REQUIRED"
	UnknownError           DaoErrorCode = "UNKNOWN_ERROR"
	DatabaseNotInitialized DaoErrorCode = "DATABASE_NOT_INITIALIZED"
)

type DaoError interface {
	error
	Code() DaoErrorCode
}

func NewDaoError(code DaoErrorCode, err error) DaoError {
	return daoError{
		error: err,
		code:  code,
	}
}

type daoError struct {
	error
	code DaoErrorCode
}

func (d daoError) Code() DaoErrorCode {
	return d.code
}
