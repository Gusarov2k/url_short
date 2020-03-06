package db_test

import (
	"context"
	"github.com/Gusarov2k/url_short"
	"github.com/Gusarov2k/url_short/internal/db"
	"testing"
)

func TestURLsRepo_Create(t *testing.T) {
	setUp(t)

	c := db.NewClient()
	if err := c.Open(PostgresTest); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = c.Close() }()

	u := shorten.URL{
		Code: "some_code",
		URL:  "http://example.org",
	}

	r := db.NewURLRepository(c)
	if err := r.Create(context.Background(), &u); err != nil {
		t.Fatal(err)
	}

	if u.ID != 1 {
		t.Fatal("bad url id, expected 1, but got: ", u.ID)
	}
}

func TestURLsRepo_ByCode(t *testing.T) {
	setUp(t)

	c := db.NewClient()
	if err := c.Open(PostgresTest); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = c.Close() }()

	u := shorten.URL{
		Code: "some_code",
		URL:  "http://example.org",
	}

	r := db.NewURLRepository(c)
	if err := r.Create(context.Background(), &u); err != nil {
		t.Fatal(err)
	}

	url, err := r.ByCode(context.Background(), "some_code")
	if err != nil {
		t.Fatal(err)
	}
	if url.Code != u.Code {
		t.Fatal("urls ids are not equal")
	}

	if url.URL != u.URL {
		t.Fatal("urls are not equal")
	}
}
