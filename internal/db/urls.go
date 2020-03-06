package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Gusarov2k/url_short"
)

const (
	msgURLInsertFailed = "Insert to url failed"
	msgURLNotFound     = "URL not found"
	msgURLScanFailed   = "Scan failed"
)

// urlsRepo is a service for managing URLs.
type urlsRepo struct {
	client *Client
}

// NewURLRepository creates a new NewURLRepository instance backed by Postgres.
func NewURLRepository(c *Client) shorten.URLRepository {
	return &urlsRepo{client: c}
}

// Create URLs's information into repository.
func (r *urlsRepo) Create(ctx context.Context, u *shorten.URL) error {
	row := r.client.db.QueryRowContext(
		ctx,
		urlInsert,
		u.Code,
		u.URL,
	)

	return shorten.WrapError(row.Scan(&u.ID), shorten.ErrInternal, msgURLInsertFailed)
}

// ByCode returns an URL object by code.
func (r *urlsRepo) ByCode(ctx context.Context, code string) (shorten.URL, error) {
	row := r.client.db.QueryRowContext(ctx, urlByCode, code)

	return scanURL(row)
}

func scanURL(rowScanner interface {
	Scan(dest ...interface{}) error
}) (shorten.URL, error) {
	var u shorten.URL
	err := rowScanner.Scan(
		&u.ID,
		&u.Code,
		&u.URL,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u, shorten.NewError(shorten.ErrURLNotFound, msgURLNotFound)
		}
		return u, shorten.WrapError(err, shorten.ErrInternal, msgURLScanFailed)
	}
	return u, nil
}
