package transfer

import (
	"context"
	"fmt"
	"time"
)

// TODO: mayble split this off into a NewViewer later on?
// would result in a bit of code duplication, I guess?

// ViewIndex returns all expenses for the given month and year and some extra information needed for displaying.
func (h transferHandler) ViewIndex(ctx context.Context) (ViewListExpensesResponse, error) {
	y, m, _ := time.Now().Date()

	res, err := h.ViewListExpenses(ctx, ListExpensesRequest{
		Year:  y,
		Month: m,
	})
	if err != nil {
		return ViewListExpensesResponse{}, fmt.Errorf("transfer.ViewIndex: %w", err)
	}

	return res, nil
}

type ViewListData struct {
	Month   time.Time
	DaySums map[int64]int64
}

type ViewListIncomesResponse struct {
	ListIncomesResponse
	ViewListData
}

// ViewListIncomes returns all incomes for the given month and year and some extra information needed for displaying.
func (h transferHandler) ViewListIncomes(ctx context.Context, req ListIncomesRequest) (ViewListIncomesResponse, error) {
	res, err := h.ListIncomes(ctx, req)
	if err != nil {
		return ViewListIncomesResponse{}, fmt.Errorf("transfer.ViewListIncomes: %w", err)
	}

	data := ViewListData{
		Month:   time.Date(req.Year, req.Month, 1, 0, 0, 0, 0, time.Local),
		DaySums: make(map[int64]int64, 31),
	}

	for _, t := range res.Incomes {
		data.DaySums[t.Date.Unix()] += t.Amount
	}

	return ViewListIncomesResponse{
		ListIncomesResponse: res,
		ViewListData:        data,
	}, nil
}

type ViewListExpensesResponse struct {
	ListExpensesResponse
	ViewListData
}

// ViewListExpenses returns all expenses for the given month and year and some extra information needed for displaying.
func (h transferHandler) ViewListExpenses(ctx context.Context, req ListExpensesRequest) (ViewListExpensesResponse, error) {
	res, err := h.ListExpenses(ctx, req)
	if err != nil {
		return ViewListExpensesResponse{}, fmt.Errorf("transfer.ViewListExpenses: %w", err)
	}

	data := ViewListData{
		Month:   time.Date(req.Year, req.Month, 1, 0, 0, 0, 0, time.Local),
		DaySums: make(map[int64]int64, 31),
	}

	for _, t := range res.Expenses {
		data.DaySums[t.Date.Unix()] += t.Amount
	}

	return ViewListExpensesResponse{
		ListExpensesResponse: res,
		ViewListData:         data,
	}, nil
}
