package httprouter

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	httputil "github.com/sollniss/kakeibo/transport/http"
	"github.com/sollniss/kakeibo/transport/http/html"
	"github.com/sollniss/kakeibo/transport/http/router"
	"github.com/sollniss/kakeibo/usecases/transfer"
)

type hrouter struct {
	*httprouter.Router
}

func NewRouter(ts router.TransferServer, views html.Views) hrouter {
	mux := httprouter.New()
	mux.HandleMethodNotAllowed = false

	mux.Handler("GET", "/", httputil.Handler(ts.ViewIndex, views.Index, views.Error))
	mux.Handler("GET", "/expenses/:year/:month", httputil.QueryHandler(urlMonthParser, ts.ViewListExpenses, views.ExpensesByMonth, views.Error))
	mux.Handler("GET", "/incomes/:year/:month", httputil.QueryHandler(urlMonthParser, ts.ViewListIncomes, views.IncomesByMonth, views.Error))

	fs, err := views.Static()
	if err != nil {
		// TODO: should we really panic here?
		panic(fmt.Errorf("httprouter.NewRouter: could not access static files: %w", err))
	}
	if fs != nil {
		mux.ServeFiles("/static/*filepath", fs)
	}

	return hrouter{mux}
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
