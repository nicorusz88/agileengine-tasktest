package controller

import (
	"github.com/agileengine-tasktest/src/api/dependencies"
	"github.com/agileengine-tasktest/src/api/errors"
	"github.com/agileengine-tasktest/src/api/model/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTransactions(c *gin.Context, container *dependencies.Container) {
	transactions := container.AccountsHandler.ListTransactions
	transactionsResponse := []vo.TransactionResponse{}
	for _, t := range transactions() {
		transactionsResponse = append(transactionsResponse, vo.TransactionResponse{
			Id:            t.Id,
			Type:          t.Type,
			Amount:        t.Amount,
			EffectiveDate: t.EffectiveDate,
		})
	}
	c.JSON(http.StatusOK, transactionsResponse)
}

func GetTransaction(c *gin.Context, container *dependencies.Container) {
	transactionId := c.Param("id")
	transaction, err := container.AccountsHandler.GetTransaction(transactionId)
	if err != nil {
		errors.ReturnError(c, err)
		return
	}
	transactionResponse := vo.TransactionResponse{
		Id:            transaction.Id,
		Type:          transaction.Type,
		Amount:        transaction.Amount,
		EffectiveDate: transaction.EffectiveDate,
	}
	c.JSON(http.StatusOK, transactionResponse)
}

func PostTransaction(c *gin.Context, container *dependencies.Container) {
	var transactionRequest vo.TransactionRequest
	err := c.ShouldBind(&transactionRequest)
	if err != nil {
		errors.ReturnError(c, &errors.Error{
			Cause:   "Invalid Arguments",
			Message: err.Error(),
		})
		return
	}

	transaction, trxError := container.AccountsHandler.CreateTransaction(transactionRequest)
	if trxError != nil {
		errors.ReturnError(c, trxError)
		return
	}
	transactionResponse := vo.TransactionResponse{
		Id:            transaction.Id,
		Type:          transaction.Type,
		Amount:        transaction.Amount,
		EffectiveDate: transaction.EffectiveDate,
	}
	c.JSON(http.StatusOK, transactionResponse)
}

func GetBalance(c *gin.Context, container *dependencies.Container) {
	balance := container.AccountsHandler.GetBalance()
	accountBalanceResponse := vo.AccountBalanceResponse{
		Balance: balance.Balance,
	}
	c.JSON(http.StatusOK, accountBalanceResponse)
}
