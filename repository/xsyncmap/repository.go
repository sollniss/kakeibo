// This implementation is whatever, just for debugging/testing.
package xsyncmap

import (
	"time"

	"github.com/puzpuzpuz/xsync/v2"
	"github.com/sollniss/kakeibo/domain"
)

type transferRepository struct {
	byMonth   *xsync.MapOf[int32, []domain.ID]
	transfers *xsync.MapOf[domain.ID, domain.Transfer]
}

func NewTransferRepository() *transferRepository {
	return &transferRepository{
		byMonth:   xsync.NewIntegerMapOf[int32, []domain.ID](),
		transfers: xsync.NewIntegerMapOf[domain.ID, domain.Transfer](),
	}
}

func (r *transferRepository) AddDummyData() {
	t1, _ := domain.NewTransfer(1, "Alice", 1000, "Salary", "Income", "2023-09-01")
	t2, _ := domain.NewTransfer(2, "Bob", 2000, "Salary", "Income", "2023-09-02")
	t3, _ := domain.NewTransfer(3, "Charlie", 3000, "Salary", "Income", "2023-09-02")
	t4, _ := domain.NewTransfer(4, "Alice", -1000, "Rent", "Rent", "2023-09-01")
	t5, _ := domain.NewTransfer(5, "Bob", -2000, "Rent", "Rent", "2023-10-01")
	t6, _ := domain.NewTransfer(6, "Charlie", 3000, "Rent", "Rent", "2023-10-01")

	r.Add(t1)
	r.Add(t2)
	r.Add(t3)
	r.Add(t4)
	r.Add(t5)
	r.Add(t6)
}

func monthKey(t time.Time) int32 {
	y, m, _ := t.Date()
	return int32(y*100 + int(m))
}

func (r *transferRepository) Add(t domain.Transfer) error {
	r.transfers.Store(t.ID, t)

	key := monthKey(t.Date)
	transfers, ok := r.byMonth.Load(key)
	if !ok {
		transfers = []domain.ID{t.ID}
	} else {
		transfers = append(transfers, t.ID)
	}

	r.byMonth.Store(key, transfers)
	return nil
}

func (r *transferRepository) Update(t domain.Transfer) error {
	r.transfers.Store(t.ID, t)
	return nil
}

func (r *transferRepository) Delete(id domain.ID) error {
	// TODO: remove from byMonth
	r.transfers.Delete(id)
	return nil
}

func (r *transferRepository) IncomesForMonth(month time.Month, year int) ([]domain.Transfer, error) {
	key := int32(year*100 + int(month))
	ids, ok := r.byMonth.Load(key)
	if !ok {
		return make([]domain.Transfer, 0), nil
	}

	transfers := make([]domain.Transfer, 0, len(ids))
	for _, id := range ids {
		t, ok := r.transfers.Load(id)
		if !ok {
			continue
		}

		if t.Amount > 0 {
			transfers = append(transfers, t)
		}
	}

	return transfers, nil
}

func (r *transferRepository) ExpensesForMonth(month time.Month, year int) ([]domain.Transfer, error) {
	key := int32(year*100 + int(month))
	ids, ok := r.byMonth.Load(key)
	if !ok {
		return make([]domain.Transfer, 0), nil
	}

	transfers := make([]domain.Transfer, 0, len(ids))
	for _, id := range ids {
		t, ok := r.transfers.Load(id)
		if !ok {
			continue
		}

		if t.Amount < 0 {
			transfers = append(transfers, t)
		}
	}

	return transfers, nil
}
