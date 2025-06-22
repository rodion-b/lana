package transactions

import "context"

type Storage interface {
	SaveTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error)
}
