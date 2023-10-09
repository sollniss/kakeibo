package transfer

import (
	"context"
	"fmt"
	"time"

	"github.com/sollniss/kakeibo/domain"
)

type transferReader interface {
	IncomesForMonth(month time.Month, year int) ([]domain.Transfer, error)
	ExpensesForMonth(month time.Month, year int) ([]domain.Transfer, error)
}

type transferWriter interface {
	Add(t domain.Transfer) error
	Update(t domain.Transfer) error
	Delete(id domain.ID) error
}

type transferHandler struct {
	idGenerator    domain.IDGenerator
	transferReader transferReader
	transferWriter transferWriter
}

func NewHandler(gen domain.IDGenerator, r transferReader, w transferWriter) transferHandler {
	return transferHandler{
		idGenerator:    gen,
		transferReader: r,
		transferWriter: w,
	}
}

type AddRequest struct {
	Person   string
	Amount   int64
	Comment  string
	Category string
	Date     string
}

type AddResponse struct {
	TransferID domain.ID
}

// Add adds a new transfer to the system.
func (h transferHandler) Add(ctx context.Context, req AddRequest) (AddResponse, error) {
	id := h.idGenerator.ID()

	t, err := domain.NewTransfer(id, req.Person, req.Amount, req.Comment, req.Category, req.Date)
	if err != nil {
		return AddResponse{}, fmt.Errorf("transfer.Add: creating transfer: %w", err)
	}

	err = h.transferWriter.Add(t)
	if err != nil {
		return AddResponse{}, fmt.Errorf("transfer.Add: adding transfer to repository: %w", err)
	}

	return AddResponse{
		TransferID: t.ID,
	}, nil
}

type UpdateRequest struct {
	TransferID domain.ID
	Person     string
	Amount     int64
	Comment    string
	Category   string
	Date       string
}

// Update updates an existing transfer in the system.
func (h transferHandler) Update(ctx context.Context, req UpdateRequest) error {
	t, err := domain.NewTransfer(req.TransferID, req.Person, req.Amount, req.Comment, req.Category, req.Date)
	if err != nil {
		return fmt.Errorf("transfer.Update: creating transfer: %w", err)
	}

	err = h.transferWriter.Update(t)
	if err != nil {
		return fmt.Errorf("transfer.Update: updating transfer in repository: %w", err)
	}

	return nil
}

type DeleteRequest struct {
	TransferID domain.ID
}

// Delete deletes an existing transfer from the system.
func (h transferHandler) Delete(ctx context.Context, req DeleteRequest) error {
	err := h.transferWriter.Delete(req.TransferID)
	if err != nil {
		return fmt.Errorf("transfer.Delete: deleting transfer in repository: %w", err)
	}

	return nil
}

type ListIncomesRequest = ByMonthRequest

type ListIncomesResponse struct {
	Incomes []domain.Transfer
}

// ListIncomes returns all incomes for the given month and year.
func (h transferHandler) ListIncomes(ctx context.Context, req ListIncomesRequest) (ListIncomesResponse, error) {
	incomes, err := h.transferReader.IncomesForMonth(req.Month, req.Year)
	if err != nil {
		return ListIncomesResponse{}, fmt.Errorf("transfer.ListIncomes: reading incomes from repository: %w", err)
	}

	return ListIncomesResponse{
		Incomes: incomes,
	}, nil
}

type ListExpensesRequest = ByMonthRequest

type ListExpensesResponse struct {
	Expenses []domain.Transfer
}

// ListExpenses returns all expenses for the given month and year.
func (u transferHandler) ListExpenses(ctx context.Context, req ListExpensesRequest) (ListExpensesResponse, error) {
	expenses, err := u.transferReader.ExpensesForMonth(req.Month, req.Year)
	if err != nil {
		return ListExpensesResponse{}, fmt.Errorf("transfer.ListExpenses: reading expenses from repository: %w", err)
	}

	return ListExpensesResponse{
		Expenses: expenses,
	}, nil
}
