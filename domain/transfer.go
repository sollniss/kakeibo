package domain

import (
	"net/http"
	"time"

	"github.com/sollniss/resperr"
)

// Transfer represents a money transfer for a user.
// The transfer can either be positive (for income) or negative (for expenses).
type Transfer struct {
	ID       ID
	Person   string
	Amount   int64
	Comment  string
	Category string // TODO: change to CategoryID
	Date     time.Time
}

func NewTransfer(id ID, person string, amount int64, comment string, cat string, date string) (Transfer, error) {
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

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return Transfer{}, resperr.WithCodeAndMessage(err, http.StatusBadRequest, "Invalid Date.")
	}

	return Transfer{
			ID:       id,
			Person:   person,
			Amount:   amount,
			Comment:  comment,
			Category: cat,
			Date:     t,
		},
		nil
}
