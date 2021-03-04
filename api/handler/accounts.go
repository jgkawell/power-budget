package handler

import (
	m "api/model"

	"github.com/gin-gonic/gin"
)

const testID = "c06d9a84-bf80-4b5e-a033-a5df8b3f1469"

// CreateAccount creates an account in the database
func (h Handler) CreateAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "CreateAccount")

	// TODO: Pull this from the request
	newAccount := m.Account{
		Name:          "Chase Debit 4",
		Balance:       4392.34,
		TotalIn:       34.55,
		TotalOut:      233.44,
		Type:          "Cash",
		CardNumber:    "****",
		AccountNumber: "****-1234",
	}

	// Call dao layer
	createdAccount, err := h.service.Accounts().CreateAccount(ctx, logger, newAccount)
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

// ReadAccount reads an account from the database with given id
func (h Handler) ReadAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "ReadAccount")

	// Call dao layer
	readAccount, err := h.service.Accounts().ReadAccount(ctx, logger, testID)
	if err != nil {
		logger.WithError(err).Error("Failed to read account")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":     "succeeded",
			"account": readAccount,
		})
	}
}

// UpdateAccount updates an account in the database with given parameters (id required)
func (h Handler) UpdateAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "UpdateAccount")

	// TODO: Pull this from the request
	newAccount := m.Account{
		ID:            testID,
		Name:          "NEW NAME",
		Balance:       3.34,
		TotalIn:       0.55,
		TotalOut:      233.44,
		Type:          "Invest",
		CardNumber:    "****",
		AccountNumber: "****-1236",
	}

	// Call dao layer
	updatedAccount, err := h.service.Accounts().UpdateAccount(ctx, logger, newAccount)
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

// DeleteAccount deletes an account in the database with given id
func (h Handler) DeleteAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "DeleteAccount")

	// Call dao layer
	deletedAccount, err := h.service.Accounts().DeleteAccount(ctx, logger, testID)
	if err != nil {
		logger.WithError(err).Error("Failed to delete account")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":     "succeeded",
			"account": deletedAccount,
		})
	}
}
