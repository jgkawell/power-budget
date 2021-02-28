package handler

import (
	"api/model"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "CreateAccount")

	// TODO: Pull this from the request
	newAccount := model.Account{
		Name:          "Chase Debit 4",
		Balance:       4392.34,
		TotalIn:       34.55,
		TotalOut:      233.44,
		Type:          "Cash",
		CardNumber:    "****",
		AccountNumber: "****-1234",
	}

	// Call dao layer
	// TODO: Need to call service layer
	createdAccount := h.db.CreateAccount(ctx, logger, newAccount)
	logger.Info(createdAccount)

	ctx.JSON(200, gin.H{
		"account": createdAccount,
	})
}

func (h Handler) ReadAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "ReadAccount")

	// Call dao layer
	// TODO: Need to call service layer
	readAccount := h.db.ReadAccount(ctx, logger, 1)
	logger.Info(readAccount)

	ctx.JSON(200, gin.H{
		"account": readAccount,
	})
}

func (h Handler) UpdateAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "UpdateAccount")

	// TODO: Pull this from the request
	newAccount := model.Account{
		ID:            3,
		Name:          "NEW NAME",
		Balance:       3.34,
		TotalIn:       0.55,
		TotalOut:      233.44,
		Type:          "Invest",
		CardNumber:    "****",
		AccountNumber: "****-1236",
	}

	// Call dao layer
	// TODO: Need to call service layer
	updatedAccount := h.db.UpdateAccount(ctx, logger, newAccount)
	logger.Info(updatedAccount)

	ctx.JSON(200, gin.H{
		"account": updatedAccount,
	})
}

func (h Handler) DeleteAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "DeleteAccount")

	// Call dao layer
	// TODO: Need to call service layer
	deletedAccount := h.db.DeleteAccount(ctx, logger, 4)
	logger.Info(deletedAccount)

	ctx.JSON(200, gin.H{
		"account": deletedAccount,
	})
}
