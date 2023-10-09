package router_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sollniss/kakeibo/transport/http/router/httprouter"
	"github.com/sollniss/kakeibo/usecases/transfer"
)

func BenchmarkHttprouter(b *testing.B) {
	router := httprouter.NewRouter(noopTransferUsecase{}, noopFrontend{})

	routes := [...]struct {
		method string
		path   string
	}{
		{"GET", "/"},
		{"GET", "/incomes/2023/09"},
		{"GET", "/incomes/2023/09/"},
		{"GET", "/expenses/2023/09"},
		{"GET", "/expenses/2023/09/"},
		{"GET", "/404"},
		{"POST", "/404"},
	}

	for _, route := range routes {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(route.method, route.path, nil)
		b.Run(route.method+" "+route.path, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				router.ServeHTTP(w, r)
			}
		})
	}
}

type noopTransferUsecase struct{}

func (noopTransferUsecase) Add(ctx context.Context, req transfer.AddRequest) (transfer.AddResponse, error) {
	return transfer.AddResponse{}, nil
}

func (noopTransferUsecase) Update(ctx context.Context, req transfer.UpdateRequest) error {
	return nil
}

func (noopTransferUsecase) Delete(ctx context.Context, req transfer.DeleteRequest) error {
	return nil
}

func (noopTransferUsecase) ListIncomes(ctx context.Context, req transfer.ListIncomesRequest) (transfer.ListIncomesResponse, error) {
	return transfer.ListIncomesResponse{}, nil
}

func (noopTransferUsecase) ListExpenses(ctx context.Context, req transfer.ListExpensesRequest) (transfer.ListExpensesResponse, error) {
	return transfer.ListExpensesResponse{}, nil
}

func (noopTransferUsecase) ViewIndex(ctx context.Context) (transfer.ViewListExpensesResponse, error) {
	return transfer.ViewListExpensesResponse{}, nil
}

func (noopTransferUsecase) ViewListIncomes(ctx context.Context, req transfer.ListIncomesRequest) (transfer.ViewListIncomesResponse, error) {
	return transfer.ViewListIncomesResponse{}, nil
}

func (noopTransferUsecase) ViewListExpenses(ctx context.Context, req transfer.ListExpensesRequest) (transfer.ViewListExpensesResponse, error) {
	return transfer.ViewListExpensesResponse{}, nil
}

type noopFrontend struct{}

func (noopFrontend) Static() (http.FileSystem, error) {
	return nil, nil
}

func (noopFrontend) ExpensesByMonth(w http.ResponseWriter, r *http.Request, res transfer.ViewListExpensesResponse) {
}

func (noopFrontend) IncomesByMonth(w http.ResponseWriter, r *http.Request, res transfer.ViewListIncomesResponse) {
}

func (noopFrontend) Index(w http.ResponseWriter, r *http.Request, resp transfer.ViewListExpensesResponse) {
}

func (noopFrontend) Error(w http.ResponseWriter, r *http.Request, err error) {
}

func (noopFrontend) NotFound(w http.ResponseWriter, r *http.Request) {
}
