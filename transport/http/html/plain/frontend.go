package plain

import (
	"net/http"
	"strconv"

	"github.com/sollniss/kakeibo/usecases/transfer"
)

type plainFrontend struct{}

func NewFrontend() plainFrontend {
	return plainFrontend{}
}

func (f plainFrontend) Index(w http.ResponseWriter, r *http.Request, res transfer.ListExpensesResponse) {
	for _, expense := range res.Expenses {
		w.Write([]byte(strconv.FormatInt(expense.Amount, 10) + "円 " + expense.Comment + "\r\n"))
	}
}

func (f plainFrontend) IncomesByMonth(w http.ResponseWriter, r *http.Request, res transfer.ListIncomesResponse) {
	for _, income := range res.Incomes {
		w.Write([]byte(strconv.FormatInt(income.Amount, 10) + "円 " + income.Comment + "\r\n"))
	}
}

func (f plainFrontend) ExpensesByMonth(w http.ResponseWriter, r *http.Request, res transfer.ListExpensesResponse) {
	for _, expense := range res.Expenses {
		w.Write([]byte(strconv.FormatInt(expense.Amount, 10) + "円 " + expense.Comment + "\r\n"))
	}
}

func (f plainFrontend) Error(w http.ResponseWriter, r *http.Request, err error) {
	w.Write([]byte(err.Error()))
}

func (f plainFrontend) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
