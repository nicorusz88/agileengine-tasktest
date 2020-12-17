package impl_test

import (
	"github.com/agileengine-tasktest/src/api/dependencies"
	"github.com/agileengine-tasktest/src/api/model/vo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddCreditTransaction(t *testing.T) {
	ass := assert.New(t)
	container := dependencies.NewMockContainer()

	amount := 100.0
	transactionRequest := vo.TransactionRequest{
		Type:   "credit",
		Amount: &amount,
	}
	createdTransaction, err := container.AccountsHandler.CreateTransaction(transactionRequest)

	ass.Nil(err)
	ass.Equal("credit", createdTransaction.Type)
	ass.Equal(100.0, createdTransaction.Amount)
	ass.NotNil(createdTransaction.Id)
	ass.NotNil(createdTransaction.EffectiveDate)
}

func TestAddDebitTransaction(t *testing.T) {
	ass := assert.New(t)
	container := dependencies.NewMockContainer()

	amount := 100.0
	transactionRequest := vo.TransactionRequest{
		Type:   "credit",
		Amount: &amount,
	}
	createdTransaction, err := container.AccountsHandler.CreateTransaction(transactionRequest)

	transactionRequest = vo.TransactionRequest{
		Type:   "debit",
		Amount: &amount,
	}
	createdTransaction, err = container.AccountsHandler.CreateTransaction(transactionRequest)

	ass.Nil(err)
	ass.Equal("debit", createdTransaction.Type)
	ass.Equal(100.0, createdTransaction.Amount)
	ass.NotNil(createdTransaction.Id)
	ass.NotNil(createdTransaction.EffectiveDate)
}

func TestAddDebitTransactionFailNegativeBalance(t *testing.T) {
	ass := assert.New(t)
	container := dependencies.NewMockContainer()

	amount := 200.0

	transactionRequest := vo.TransactionRequest{
		Type:   "debit",
		Amount: &amount,
	}
	_, err := container.AccountsHandler.CreateTransaction(transactionRequest)

	ass.NotNil(err)
	ass.Equal(err.Cause, "NEGATIVE_BALANCE")
	ass.Equal(err.Message, "Account balance cannot be negative")
	ass.Equal(err.Code.Status, 422)
}

func TestGetBalance(t *testing.T) {
	ass := assert.New(t)
	container := dependencies.NewMockContainer()

	amount := 200.0
	transactionRequest := vo.TransactionRequest{
		Type:   "credit",
		Amount: &amount,
	}
	container.AccountsHandler.CreateTransaction(transactionRequest)

	amount = 100.0
	transactionRequest = vo.TransactionRequest{
		Type:   "debit",
		Amount: &amount,
	}
	container.AccountsHandler.CreateTransaction(transactionRequest)

	balance := container.AccountsHandler.GetBalance()

	ass.Equal(200.0, balance.Balance)
}

func TestListTransactions(t *testing.T) {
	ass := assert.New(t)
	container := dependencies.NewMockContainer()

	transactions := container.AccountsHandler.ListTransactions()
	if len(transactions) != 0 {
		transaction := transactions[0]
		ass.NotNil(transaction)
		ass.NotNil(transaction.Id)
		ass.NotNil(transaction.Amount)
		ass.NotNil(transaction.Type)
		ass.NotNil(transaction.EffectiveDate)
	}
}

func TestGetTransaction(t *testing.T) {
	ass := assert.New(t)
	container := dependencies.NewMockContainer()

	transactions := container.AccountsHandler.ListTransactions()
	if len(transactions) != 0 {
		transaction := transactions[0]
		getTransaction, err := container.AccountsHandler.GetTransaction(transaction.Id)
		ass.Nil(err)
		ass.NotNil(getTransaction)
		ass.NotNil(getTransaction.Id)
		ass.NotNil(getTransaction.Amount)
		ass.NotNil(getTransaction.Type)
		ass.NotNil(getTransaction.EffectiveDate)

	}
}

