package quicktemplate

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/sollniss/kakeibo/transport/http/html/quicktemplate/templates"
	"github.com/sollniss/kakeibo/usecases/transfer"
)

type tpl struct{}

func NewFrontend() tpl {
	return tpl{}
}

//go:embed static/*
var staticContent embed.FS

func (t tpl) Static() (http.FileSystem, error) {
	serverRoot, err := fs.Sub(staticContent, "static")
	if err != nil {
		return nil, err
	}
	return http.FS(serverRoot), nil
}

func (t tpl) Index(w http.ResponseWriter, r *http.Request, res transfer.ViewListExpensesResponse) {
	t.ExpensesByMonth(w, r, res)
}

func (t tpl) IncomesByMonth(w http.ResponseWriter, r *http.Request, res transfer.ViewListIncomesResponse) {
	p := &templates.TransferList{
		Base: templates.Base{
			SubTitle: "今月",
		},
		Transfers: res.Incomes,
		Current:   "incomes",
		Data:      res.ViewListData,
	}

	templates.WritePage(w, p)
}

func (t tpl) ExpensesByMonth(w http.ResponseWriter, r *http.Request, res transfer.ViewListExpensesResponse) {
	p := &templates.TransferList{
		Base: templates.Base{
			SubTitle: "今月",
		},
		Transfers: res.Expenses,
		Current:   "expenses",
		Data:      res.ViewListData,
	}

	templates.WritePage(w, p)
}

func (t tpl) Error(w http.ResponseWriter, r *http.Request, err error) {
	w.Write([]byte(err.Error()))
}

func (t tpl) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
