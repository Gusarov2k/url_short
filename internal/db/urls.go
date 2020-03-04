package db

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Url struct {
	ID   int    `json:"id"`
	Url  string `json:"url"`
	Code string `json:"code"`
}

type UrlRepository interface {
	Create(ctx context.Context, u Url) (int64, error)
}

func ConnToDB(Conn *sqlx.DB) UrlRepository {
	return &Client{
		db: Conn,
	}
}

func (con *Client) CreateTableIsfExist(schema string) {
	con.db.MustExec(schema)
}

func (con *Client) Create(ctx context.Context, u Url) (int64, error) {
	query := "Insert urls SET url=?, code=?"

	stmt, error := con.db.PrepareContext(ctx, query)
	if error != nil {
		return -1, error
	}

	res, error := stmt.ExecContext(ctx, u.Url, u.Code)

	if error != nil {
		return -1, error
	}

	return res.LastInsertId()
}
