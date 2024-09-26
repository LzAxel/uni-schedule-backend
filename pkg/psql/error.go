package psql

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

func IsPgErrorCode(err error, code pq.ErrorCode) bool {
	var pgErr *pq.Error
	ok := errors.As(err, &pgErr)
	if ok && pgErr.Code == code {
		return true
	}

	return false
}

func IsNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
