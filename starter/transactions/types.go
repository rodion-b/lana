package transactions

import "time"

type Money struct {
	Amount   int
	Currency string
}

func NewMoney(amount int, currency string) Money {
	return Money{Amount: amount, Currency: currency}
}

func GBP(amount int) Money {
	return NewMoney(amount, "GBP")
}

func KES(amount int) Money {
	return NewMoney(amount, "KES")
}

type Transaction struct {
	UserID         string
	AmountSent     Money
	AmountReceived Money
	Created        time.Time
}

func NewTransaction(userID string, amountSent Money, amountReceived Money) (*Transaction, error) {
	return &Transaction{
		UserID:         userID,
		AmountSent:     amountSent,
		AmountReceived: amountReceived,
		Created:        time.Now(),
	}, nil
}
