package db

import (
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	defaultMaxConnections = 5
)

type Client struct {
	db             *sqlx.DB
	logger         log.Logger
	maxConnections int
}

// Open connection to PostgreSQL.
func (c *Client) Open(dataSourceName string) error {
	var err error

	c.logger.Log("level", "debug", "msg", "connecting to db")
	if c.db, err = sqlx.Open("postgres", dataSourceName); err != nil {
		c.logger.Log("level", "debug", "msg", "test")
		return err
	}
	if err = c.db.Ping(); err != nil {
		return err
	}
	c.db.SetMaxOpenConns(c.maxConnections)
	c.db.SetMaxIdleConns(c.maxConnections)
	c.logger.Log("level", "debug", "msg", "connected to db")

	return nil
}

// Close closes PostgreSQL connection.
func (c *Client) Close() error {
	return c.db.Close()
}
