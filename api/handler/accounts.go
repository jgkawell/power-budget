package handler

import (
	"net/http"

	m "api/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateAccount creates an account in the database
func (h Handler) CreateAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "CreateAccount")

	// Pull out account from request
	var newAccount m.Account
	if err := ctx.ShouldBindJSON(&newAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call dao layer
	account, err := h.service.Accounts().CreateAccount(ctx, logger, newAccount)
	if err != nil {
		logger.WithError(err).Error("Failed to create account")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, account)
	}
}

// ReadAccount reads an account from the database with given id
func (h Handler) ReadAccount(ctx *gin.Context) {
	// Pull out id to read and add to logger
	id := ctx.Param("id")
	logger := h.logger.WithFields(logrus.Fields{
		"account_id": id,
		"handler":    "ReadAccount",
	})

	// Call dao layer
	account, err := h.service.Accounts().ReadAccount(ctx, logger, id)
	if err != nil {
		logger.WithError(err).Error("Failed to read account")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, account)
	}
}

// ReadAllAccounts reads all accounts from the database
func (h Handler) ReadAllAccounts(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "ReadAllAccounts")

	// Call dao layer
	accounts, err := h.service.Accounts().ReadAllAccounts(ctx, logger)
	if err != nil {
		logger.WithError(err).Error("Failed to read all accounts")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, accounts)
	}
}

// UpdateAccount updates an account in the database with given parameters (id required)
func (h Handler) UpdateAccount(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "UpdateAccount")

	// Pull out account from request
	var newAccount m.Account
	if err := ctx.ShouldBindJSON(&newAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger = logger.WithField("account_id", newAccount.ID)

	// Call dao layer
	account, err := h.service.Accounts().UpdateAccount(ctx, logger, newAccount)
	if err != nil {
		logger.WithError(err).Error("Failed to update account")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, account)
	}
}

// DeleteAccount deletes an account in the database with given id
func (h Handler) DeleteAccount(ctx *gin.Context) {
	// Pull out id to delete and add to logger
	id := ctx.Param("id")
	logger := h.logger.WithFields(logrus.Fields{
		"account_id": id,
		"handler":    "DeleteAccount",
	})

	// Call dao layer
	account, err := h.service.Accounts().DeleteAccount(ctx, logger, id)
	if err != nil {
		logger.WithError(err).Error("Failed to delete account")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, account)
	}
}
