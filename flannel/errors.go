package flannel

import (
	"errors"
)

var (
	// ErrRecordNotFound returns a "record not found error". Occurs only when attempting to query the database with a struct; querying with a slice won't return this error
	ErrRecordNotFound   = errors.New("record not found")
	ErrExistKeyInLedger = errors.New("ledger is already exist this record")
)

// IsRecordNotFoundError returns true if error contains a RecordNotFound error
func IsRecordNotFoundError(err error) bool {
	return err == ErrRecordNotFound
}

func IsExistKeyInLedgerError(err error) bool {
	return err == ErrExistKeyInLedger
}
