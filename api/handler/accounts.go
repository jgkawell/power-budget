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
	createdAccount, err := h.db.CreateAccount(ctx, logger, newAccount)
	if err != nil {
		logger.WithError(err).Error("Failed to create account")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":     "succeeded",
			"account": createdAccount,
		})
	}
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
	updatedAccount, err := h.db.UpdateAccount(ctx, logger, newAccount)
	if err != nil {
		logger.WithError(err).Error("Failed to update account")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":     "succeeded",
			"account": updatedAccount,
		})
	}
}

func (h Handler) DeleteAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "DeleteAccount")

	// Call dao layer
	// TODO: Need to call service layer
	deletedAccount := h.db.DeleteAccount(ctx, logger, "dbc23cf1-6503-4996-bbe0-ffb569995639")
	logger.Info(deletedAccount)

	ctx.JSON(200, gin.H{
		"msg": deletedAccount,
	})
}
