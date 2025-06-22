package main

import (
	"context"
	"fmt"
	"log"

	"github.com/NalaMoney/interviews/go/transactions"
)

func main() {
	storage, err := transactions.NewMemoryStorage()
	if err != nil {
		log.Fatalln("could not initialise storage")
	}

	transactionsAPI, err := transactions.NewAPI(storage)
	if err != nil {
		log.Fatal("could not initialise API")
	}

	ctx := context.Background()

	transaction, err := transactionsAPI.CreateTransaction(ctx, "123", transactions.GBP(100), transactions.KES(15000))
	if err != nil {
		log.Fatalln("could not create transaction")
	}

	fmt.Println(*transaction) // nolint
}
