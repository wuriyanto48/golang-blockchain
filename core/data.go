package core

import (
	"time"
)

// Transaction simple models
type Transaction struct {
	From   string
	To     string
	Amount float64
	Date   time.Time
}

// NewTransaction function, Transaction's constructor
func NewTransaction(from, to string, amount float64) *Transaction {
	return &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
		Date:   time.Now(),
	}
}
