package shorten

import "context"

type URL struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Code string `json:"code"`
}

// UrlRepository is a storage for urls.
type URLRepository interface {
	// Create creates a new url.
	Create(ctx context.Context, u *URL) error
	// ByCode retrieves an url by its code.
	ByCode(ctx context.Context, code string) (URL, error)
}
