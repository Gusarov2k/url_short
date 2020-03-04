package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var client Client

func TestLinkDB(t *testing.T) {
	err := client.Open("host=localhost port=5432 user=ivan dbname=short_link_test password=1234 sslmode=disable")

	assert.Nil(t, err, nil)

}

func TestCloseLinkDB(t *testing.T) {
	client.Open("host=localhost port=5432 user=ivan dbname=short_link_test password=1234 sslmode=disable")
	err := client.Close()
	assert.Nil(t, err, nil)
}

func TestCreateTables(t *testing.T) {

	client.Open("host=localhost port=5432 user=ivan dbname=short_link_test password=1234 sslmode=disable")
	client.CreateTableIsfExist(Schema)
	// assert.Nil(t, err, nil)
}

func TestCreateUrl(t *testing.T) {
	var url = Url{Url: "test", Code: "code test"}
	client.Open("host=localhost port=5432 user=ivan dbname=short_link_test password=1234 sslmode=disable")
	client.Create(url)
	// assert.Nil(t, err, nil)
}
