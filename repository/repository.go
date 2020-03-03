package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"url_short/connectdb"
)

type Url struct {
	ID   int    `json:"id"`
	Url  string `json:"url"`
	Code string `json:"code"`
}

type UrlRepository interface {
	CreateTable(query string) error
	Create(ctx context.Context, u Url) error
}

func CreateTableIfExist(urlRepository UrlRepository) {
	urlRepository.CreateTable()
}

func (u Url) CreateTableIfExist(schema string, con db.Client) error {
	_, error := con.MustExec(schema)
	return error
}
