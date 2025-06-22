package transactions

import "context"

type MemoryStorage struct {
	Transactions []*Transaction
}

func NewMemoryStorage() (*MemoryStorage, error) {
	return &MemoryStorage{
		Transactions: make([]*Transaction, 0, 100),
	}, nil
}

func (m *MemoryStorage) SaveTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error) {
	m.Transactions = append(m.Transactions, transaction)
	return transaction, nil
}
