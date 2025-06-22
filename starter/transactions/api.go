package transactions

import "context"

type API struct {
	storage Storage
}

func NewAPI(storage Storage) (*API, error) {
	return &API{
		storage: storage,
	}, nil
}

func (a *API) CreateTransaction(ctx context.Context, userID string, amountSent Money, amountReceived Money) (*Transaction, error) {
	transaction, err := NewTransaction(userID, amountSent, amountReceived)
	if err != nil {
		return nil, err
	}

	transaction, err = a.storage.SaveTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
