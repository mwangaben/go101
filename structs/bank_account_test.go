package structs

import (
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestNewBankAccount(t *testing.T) {
	account := NewBankAccount("Benedict", 0)

	assert2.Equal(t, 0, account.balance)

	account.Deposit(2000)
	assert2.Equal(t, 2000, account.balance)
	account.Withdrawal(200)
	assert2.Equal(t, 1800, account.balance)

	account.Deposit(2000)
	account.Withdrawal(500)
	assert2.Equal(t, 3300, account.GetBalance())

	assert2.Equal(t, 2000, account.GetTransactions()[0].Amount)
	assert2.Equal(t, 4, len(account.GetTransactions()))

	//fmt.Printf("the transactions: %v \n", account.GetTransactions())
}
