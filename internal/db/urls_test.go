package db

import (
	"github.com/stretchr/testify"
	"testing"
)

func TestLinkDB(t *testing.T) {
	var client Client
	db := client.Open("host=localhost port=5432 user=ivan dbname=short_link_development password=1234 sslmode=disable")
	var url Url
	// assert for nil (good for errors)
	assert.Nil(t, url.CreateTableIfExist(Schema, db), error)

}
