module github.com/Gusarov2k/url_short.git

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-playground/validator/v10 v10.2.0
	github.com/lib/pq v1.1.1
	url_short/connectdb v0.0.1
)

replace url_short/connectdb v0.0.1 => ./internal/db
