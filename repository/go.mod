module url_short/repository

require url_short/connectdb v0.0.1
replace url_short/connectdb v0.0.1 => ../internal/db
go 1.13
