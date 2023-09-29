package domain

import (
	"net/http"

	"github.com/sollniss/resperr"
)

// Transfer represents a money transfer for a user.
// The transfer can either be positive (for income) or negative (for expenses).
type Transfer struct {
	TransferID ID
	Person     string
	Amount     int64
	Comment    string
	Category   string // TODO: change to CategoryID
}

func NewTransfer(id ID, person string, amount int64, comment string, cat string) (Transfer, error) {
	// Validate the input
	if amount == 0 {
		return Transfer{}, resperr.WithCodeAndMessage(nil, http.StatusBadRequest, "Invalid Amount (0).")
	}

	if person == "" {
		return Transfer{}, resperr.WithCodeAndMessage(nil, http.StatusBadRequest, "Person must not be empty.")
	}

	if cat == "" {
		return Transfer{}, resperr.WithCodeAndMessage(nil, http.StatusBadRequest, "Category must not be empty.")
	}

	return Transfer{
			TransferID: id,
			Person:     person,
			Amount:     amount,
			Comment:    comment,
			Category:   cat,
		},
		nil
}
