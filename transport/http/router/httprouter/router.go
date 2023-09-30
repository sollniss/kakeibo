package httprouter

import (
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	httputil "github.com/sollniss/kakeibo/transport/http"
	"github.com/sollniss/kakeibo/transport/http/html"
	"github.com/sollniss/kakeibo/transport/http/router"
	"github.com/sollniss/kakeibo/usecases/transfer"
)

func NewRouter(ts router.TransferServer, views html.Views) *httprouter.Router {
	mux := httprouter.New()
	mux.HandleMethodNotAllowed = false

	mux.Handler("GET", "/expenses/:year/:month", httputil.QueryHandler(urlMonthParser, ts.ListExpenses, views.ExpensesByMonth, views.Error))
	mux.Handler("GET", "/incomes/:year/:month", httputil.QueryHandler(urlMonthParser, ts.ListIncomes, views.IncomesByMonth, views.Error))

	return mux
}

func urlMonthParser(w http.ResponseWriter, r *http.Request) (transfer.ByMonthRequest, error) {
	var req transfer.ByMonthRequest
	params := httprouter.ParamsFromContext(r.Context())

	month, err := strconv.Atoi(params.ByName("month"))
	if err != nil {
		return req, err
	}

	year, err := strconv.Atoi(params.ByName("year"))
	if err != nil {
		return req, err
	}

	req.Month = time.Month(month)
	req.Year = year

	return req, nil
}
