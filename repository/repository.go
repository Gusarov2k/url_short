package repository

import (
	"context"
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

// func CreateUrl(urlRepository UrlRepository) {
// 	urlRepository.Create()
// }

func CreateTableIfExist(urlRepository UrlRepository) {
	urlRepository.CreateTable()
}

// func (u Url) CreateUrl(ctx context.Context, u Url) error {
// 	query := "Insert urls SET url=?, code=?"

// 	stmt, err := m.Conn.PrepareContext(ctx, query)
// 	if err != nil {
// 		return -1, err
// 	}

// 	defer stmt.Close()

// 	if err != nil {
// 		return -1, err
// 	}

// 	return res.LastInsertId()

// }

func (u Url) CreateTableIfExist(schema string, db Client) error {
	rows, error := db.MustExec(schema)
	return error
}
