package db

import (
	// "context"
	"github.com/jmoiron/sqlx"
)

type Url struct {
	ID   int    `json:"id"`
	Url  string `json:"url"`
	Code string `json:"code"`
}

// type UrlRepository interface {
// 	CreateTable(query string) error
// }

// func CreateTableIfExist(urlRepository UrlRepository) {
// 	urlRepository.CreateTable()
// }

func (u Url) CreateTableIfExist(schema string, con sqlx.DB) error {
	error := con.MustExec(schema)
	return error
}
