package handler

import (
	"time"

	m "api/model"

	"github.com/gin-gonic/gin"
)

var testCreditID = "8e91655c-805b-4b08-8f29-09f7cc3d885a"

const testCreditAccountID = "c06d9a84-bf80-4b5e-a033-a5df8b3f1468"

// CreateCredit creates a credit in the database
func (h Handler) CreateCredit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "CreateCredit")

	// TODO: Pull this from the request
	newCredit := m.Credit{
		PostedDate: time.Now(),
		Amount:     92.34,
		Source:     "Chase Debit",
		Purpose:    "Groceries",
		AccountID:  testCreditAccountID,
		Budget:     2,
		Notes:      "Big refund",
	}

	// Call dao layer
	createdCredit, err := h.service.Credits().CreateCredit(ctx, logger, newCredit)
	if err != nil {
		logger.WithError(err).Error("Failed to create credit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		testCreditID = createdCredit.ID
		ctx.JSON(200, gin.H{
			"msg":    "succeeded",
			"credit": createdCredit,
		})
	}
}

// ReadCredit reads a credit from the database with given id
func (h Handler) ReadCredit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "ReadCredit")

	// Call dao layer
	readCredit, err := h.service.Credits().ReadCredit(ctx, logger, testCreditID)
	if err != nil {
		logger.WithError(err).Error("Failed to read credit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":    "succeeded",
			"credit": readCredit,
		})
	}
}

// UpdateCredit updates a credit in the database with given parameters (id required)
func (h Handler) UpdateCredit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "UpdateCredit")

	// TODO: Pull this from the request
	newCredit := m.Credit{
		ID:         testCreditID,
		PostedDate: time.Now(),
		Amount:     92.34,
		Source:     "NEW SOURCE",
		Purpose:    "Groceries",
		AccountID:  testCreditAccountID,
		Budget:     2,
		Notes:      "Big purchase",
	}

	// Call dao layer
	updatedCredit, err := h.service.Credits().UpdateCredit(ctx, logger, newCredit)
	if err != nil {
		logger.WithError(err).Error("Failed to update credit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":    "succeeded",
			"credit": updatedCredit,
		})
	}
}

// DeleteCredit deletes a credit in the database with given id
func (h Handler) DeleteCredit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "DeleteCredit")

	// Call dao layer
	deletedCredit, err := h.service.Credits().DeleteCredit(ctx, logger, testCreditID)
	if err != nil {
		logger.WithError(err).Error("Failed to delete credit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":    "succeeded",
			"credit": deletedCredit,
		})
	}
}
