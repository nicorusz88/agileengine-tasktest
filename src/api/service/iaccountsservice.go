package service

import (
	"github.com/agileengine-tasktest/src/api/errors"
	"github.com/agileengine-tasktest/src/api/model/account"
	"github.com/agileengine-tasktest/src/api/model/vo"
	"github.com/agileengine-tasktest/src/api/repository"
	"github.com/agileengine-tasktest/src/api/service/impl"
)

type AccountsHandler interface {
	CreateTransaction(transactionRequest vo.TransactionRequest) (account.Transaction, *errors.Error)
	ListTransactions() []account.Transaction
	GetBalance() account.AccountBalance
	GetTransaction(transactionId string) (account.Transaction, *errors.Error)
}

func NewAccountsHandler() AccountsHandler {
	return &impl.AccountTransactionService{TransactionRepository: repository.NewAccountsPersistenceHandler()}
}
