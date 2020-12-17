package repository

import (
	"github.com/agileengine-tasktest/src/api/errors"
	"github.com/agileengine-tasktest/src/api/model/account"
	repository "github.com/agileengine-tasktest/src/api/repository/impl"
)

type AccountsPersistenceHandler interface{
	PersistTransaction(account.Transaction) (account.Transaction, *errors.Error)
	RetrieveTransactions() ([]account.Transaction)
	RetrieveBalance() float64
	RetrieveTransaction(transactionId string) (account.Transaction, *errors.Error)
}

func NewAccountsPersistenceHandler() AccountsPersistenceHandler {
	return &repository.MemoryTransactionPersistenceHandler{}
}

