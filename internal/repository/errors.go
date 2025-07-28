package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	Error             = errors.New("something went wrong")
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicatedKey  = errors.New("duplicated key")
)

func isDuplicateKeyError(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

func SqlcErrToRepositoryErr(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return ErrRecordNotFound
	}

	if isDuplicateKeyError(err) {
		return ErrDuplicatedKey
	}

	slog.Error(fmt.Sprintf("Register unprocessable error: %v", err))
	return Error
}
