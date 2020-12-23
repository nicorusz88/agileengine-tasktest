package repository

import (
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
	balanceMutex   = &sync.RWMutex{}
)

func (*MemoryTransactionPersistenceHandler) PersistTransaction(transaction account.Transaction) (account.Transaction, *errors.Error) {
	balanceMutex.Lock()
	defer balanceMutex.Unlock()
	if transaction.Type == string(CREDIT) {
		currentBalance = currentBalance + transaction.Amount
	} else {
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
	balanceMutex.RLock()
	defer balanceMutex.RUnlock()
	return currentBalance
}
