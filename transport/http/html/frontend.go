package html

import (
	"net/http"

	"github.com/sollniss/kakeibo/usecases/transfer"
)

// Views is a set of views for the frontend.
// The functions write to the ResponseWriter so no errors are returned.
type Views interface {
	Index(w http.ResponseWriter, r *http.Request, res transfer.ListExpensesResponse)
	IncomesByMonth(w http.ResponseWriter, r *http.Request, res transfer.ListIncomesResponse)
	ExpensesByMonth(w http.ResponseWriter, r *http.Request, res transfer.ListExpensesResponse)
	Error(w http.ResponseWriter, r *http.Request, err error)
	NotFound(w http.ResponseWriter, r *http.Request)
}
