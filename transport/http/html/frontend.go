package html

import (
	"net/http"

	"github.com/sollniss/kakeibo/transport/http/html/quicktemplate"
	"github.com/sollniss/kakeibo/usecases/transfer"
)

var (
	_ = Views(quicktemplate.NewFrontend())
	//_ = Views(plain.NewFrontend())
)

// Views is a set of views for the frontend.
// The functions write to the ResponseWriter so no errors are returned.
type Views interface {
	// Static returns statically served files. Might return nil if no static files need to be served.
	Static() (http.FileSystem, error)
	Index(w http.ResponseWriter, r *http.Request, res transfer.ViewListExpensesResponse)
	IncomesByMonth(w http.ResponseWriter, r *http.Request, res transfer.ViewListIncomesResponse)
	ExpensesByMonth(w http.ResponseWriter, r *http.Request, res transfer.ViewListExpensesResponse)
	Error(w http.ResponseWriter, r *http.Request, err error)
	NotFound(w http.ResponseWriter, r *http.Request)
}
