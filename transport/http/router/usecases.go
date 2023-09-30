package router

import (
	"context"

	"github.com/sollniss/kakeibo/usecases/transfer"
)

type TransferServer interface {
	Add(ctx context.Context, req transfer.AddRequest) (transfer.AddResponse, error)
	Update(ctx context.Context, req transfer.UpdateRequest) error
	Delete(ctx context.Context, req transfer.DeleteRequest) error
	ListIncomes(ctx context.Context, req transfer.ListIncomesRequest) (transfer.ListIncomesResponse, error)
	ListExpenses(ctx context.Context, req transfer.ListExpensesRequest) (transfer.ListExpensesResponse, error)
}
