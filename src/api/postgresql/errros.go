package postgresql

import "errors"

var (
	TransactionNotFound = errors.New("transaction not found")
)
