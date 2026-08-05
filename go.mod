module github.com/storage-lock/go-us3-storage

go 1.19

require (
	github.com/storage-lock/go-sql-based-storage v0.0.2
	github.com/storage-lock/go-storage v0.0.2
	github.com/storage-lock/go-storage-lock v0.0.3
	github.com/storage-lock/go-storage-test-helper v0.0.3
	github.com/stretchr/testify v1.8.4
)

replace github.com/storage-lock/go-storage => ../go-storage
replace github.com/storage-lock/go-sql-based-storage => ../go-sql-based-storage
