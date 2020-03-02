package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Client struct {
	db             *sql.DB
	logger         log.Logger
	maxConnections int
}

// Open connection to PostgreSQL.
func (c *Client) Open(dataSourceName string) error {
	db, err := sql.Open("postgres", dataSourceName)

	if c.maxConnections > 0 {
		db.SetMaxOpenConns(c.maxConnections)
	}
	c.db = db
}

// Close closes PostgreSQL connection.
func (c *Client) Close() error {
	c.db.Close()
}
