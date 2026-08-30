package structs

import "time"

type BankAccount struct {
	Owner        string
	balance      int
	Transactions []Transaction
}

func NewBankAccount(owner string, balance int) *BankAccount {
	return &BankAccount{
		Owner:        owner,
		balance:      balance,
		Transactions: []Transaction{},
	}
}

func (ba *BankAccount) Deposit(amount int) {
	ba.balance += amount
	ba.Transactions = append(ba.Transactions, Transaction{
		Amount: amount,
		Type:   "deposit",
		Date:   time.Now(),
	})
}

func (ba *BankAccount) Withdrawal(amount int) {
	ba.balance -= amount
	ba.Transactions = append(ba.Transactions, Transaction{
		Amount: amount,
		Type:   "withdrawal",
		Date:   time.Now(),
	})
}

func (ba *BankAccount) GetTransactions() []Transaction {
	return ba.Transactions
}

func (ba *BankAccount) GetBalance() int {
	return ba.balance
}

type Transaction struct {
	Amount int
	Type   string
	Date   time.Time
}
