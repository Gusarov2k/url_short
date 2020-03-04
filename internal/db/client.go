package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Client struct {
	db             *sqlx.DB
	logger         *log.Logger
	maxConnections int
}

// Open connection to PostgreSQL.
func (c *Client) Open(dataSourceName string) error {
	db, error := sqlx.Open("postgres", dataSourceName)
	if c.maxConnections > 0 {
		db.SetMaxOpenConns(c.maxConnections)
	}

	// c.logger = log.New(db, "INFO: ", log.Ldate|log.Lshortfile)
	c.db = db

	return error
}

// Close closes PostgreSQL connection.
func (c *Client) Close() error {
	error := c.db.Close()
	return error
}
