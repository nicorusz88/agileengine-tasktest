package repository

import (
	"fmt"
	"github.com/agileengine-tasktest/src/api/errors"
	"github.com/agileengine-tasktest/src/api/model/account"
	"sync"
)

type MemoryTransactionPersistenceHandler struct{}

func (*MemoryTransactionPersistenceHandler) RetrieveTransaction(transactionId string) (account.Transaction, *errors.Error) {
	transaction, ok := transactions.Load(transactionId)
	if !ok {
		err := &errors.Error{
			Message: "Inexistent transaction",
			Code:    errors.NotFoundApiError,
			Cause:   "TRANSACTION_NOT_FOUND",
		}
		return account.Transaction{}, err
	}
	return transaction.(account.Transaction), nil
}

type TransactionTypeEnum string

const (
	DEBIT  TransactionTypeEnum = "debit"
	CREDIT TransactionTypeEnum = "credit"
)

var (
	transactions   = sync.Map{}
	currentBalance float64
	balanceMutex   = &sync.Mutex{}
)

func (*MemoryTransactionPersistenceHandler) PersistTransaction(transaction account.Transaction) (account.Transaction, *errors.Error) {
	if transaction.Type == string(CREDIT) {
		balanceMutex.Lock()
		defer balanceMutex.Unlock()
		currentBalance = currentBalance + transaction.Amount
	} else {
		balanceMutex.Lock()
		defer balanceMutex.Unlock()
		currentBalanceAux := currentBalance - transaction.Amount
		if currentBalanceAux < 0 {
			err := &errors.Error{
				Message: "Account balance cannot be negative",
				Code:    errors.UnprocessableEntityApiError,
				Cause:   "NEGATIVE_BALANCE",
			}
			return transaction, err
		}
		currentBalance = currentBalanceAux
		fmt.Printf("Account Balance %f", currentBalance)
	}
	transactions.Store(transaction.Id, transaction)
	return transaction, nil
}

func (*MemoryTransactionPersistenceHandler) RetrieveTransactions() []account.Transaction {
	transactionsSlice := []account.Transaction{}
	transactions.Range(func(key, value interface{}) bool {
		transactionsSlice = append(transactionsSlice, value.(account.Transaction))
		return true
	})
	return transactionsSlice
}

func (*MemoryTransactionPersistenceHandler) RetrieveBalance() float64 {
	return currentBalance
}
