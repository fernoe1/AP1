package Model

import (
	"errors"
	"strconv"
)

type BankAccount struct {
	Balance      float64
	Transactions []string
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("negative amount")
	}

	b.Balance += amount
	b.Transactions = append(b.Transactions, "Deposited "+strconv.FormatFloat(amount, 'f', -1, 64))

	return nil
}

func (b *BankAccount) Withdraw(amount float64) (float64, error) {
	if amount < 0 {
		return 0.0, errors.New("negative amount")
	}

	b.Balance -= amount
	b.Transactions = append(b.Transactions, "Withdrew "+strconv.FormatFloat(amount, 'f', -1, 64))

	return b.Balance, nil
}

func (b *BankAccount) GetBalance() (float64, []string) {
	return b.Balance, b.Transactions
}
