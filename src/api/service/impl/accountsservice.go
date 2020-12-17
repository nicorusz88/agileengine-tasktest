package impl

import (
	"fmt"
	"github.com/agileengine-tasktest/src/api/errors"
	"github.com/agileengine-tasktest/src/api/model/account"
	"github.com/agileengine-tasktest/src/api/model/vo"
	"github.com/agileengine-tasktest/src/api/repository"
	"github.com/agileengine-tasktest/src/api/utils"
	"time"
)

type AccountTransactionService struct {
	TransactionRepository repository.AccountsPersistenceHandler
}

func (tr *AccountTransactionService) GetTransaction(transactionId string) (account.Transaction, *errors.Error) {
	return tr.TransactionRepository.RetrieveTransaction(transactionId)
}

func (tr *AccountTransactionService) GetBalance() account.AccountBalance {
	balance := tr.TransactionRepository.RetrieveBalance()

	return account.AccountBalance{
		Balance: balance,
	}
}

func (tr *AccountTransactionService) CreateTransaction(transactionRequest vo.TransactionRequest) (account.Transaction, *errors.Error) {

	transactionToPersist := account.Transaction{
		Id:            fmt.Sprintf("%s", utils.GetNewUUID()),
		Type:          transactionRequest.Type,
		Amount:        *transactionRequest.Amount,
		EffectiveDate: time.Now(),
	}

	return tr.TransactionRepository.PersistTransaction(transactionToPersist)
}

func (tr *AccountTransactionService) ListTransactions() []account.Transaction {
	return tr.TransactionRepository.RetrieveTransactions()
}
