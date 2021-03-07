package handler

import (
	"time"

	m "api/model"

	"github.com/gin-gonic/gin"
)

var testDebitID = "c06d9a84-bf80-4b5e-a033-a5df8b3f1469"

const testDebitAccountID = "c06d9a84-bf80-4b5e-a033-a5df8b3f1468"

// CreateDebit creates a debit in the database
func (h Handler) CreateDebit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "CreateDebit")

	// TODO: Pull this from the request
	newDebit := m.Debit{
		PostedDate: time.Now(),
		Amount:     92.34,
		Vendor:     "Walmart",
		Purpose:    "Groceries",
		AccountID:  testDebitAccountID,
		Budget:     2,
		Notes:      "Big purchase",
	}

	// Call dao layer
	createdDebit, err := h.service.Debits().CreateDebit(ctx, logger, newDebit)
	if err != nil {
		logger.WithError(err).Error("Failed to create debit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		testDebitID = createdDebit.ID
		ctx.JSON(200, gin.H{
			"msg":   "succeeded",
			"debit": createdDebit,
		})
	}
}

// ReadDebit reads a debit from the database with given id
func (h Handler) ReadDebit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "ReadDebit")

	// Call dao layer
	readDebit, err := h.service.Debits().ReadDebit(ctx, logger, testDebitID)
	if err != nil {
		logger.WithError(err).Error("Failed to read debit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":   "succeeded",
			"debit": readDebit,
		})
	}
}

// UpdateDebit updates a debit in the database with given parameters (id required)
func (h Handler) UpdateDebit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "UpdateDebit")

	// TODO: Pull this from the request
	newDebit := m.Debit{
		ID:         testDebitID,
		PostedDate: time.Now(),
		Amount:     92.34,
		Vendor:     "NEW VENDOR",
		Purpose:    "Groceries",
		AccountID:  testDebitAccountID,
		Budget:     2,
		Notes:      "Big purchase",
	}

	// Call dao layer
	updatedDebit, err := h.service.Debits().UpdateDebit(ctx, logger, newDebit)
	if err != nil {
		logger.WithError(err).Error("Failed to update debit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":   "succeeded",
			"debit": updatedDebit,
		})
	}
}

// DeleteDebit deletes a debit in the database with given id
func (h Handler) DeleteDebit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "DeleteDebit")

	// Call dao layer
	deletedDebit, err := h.service.Debits().DeleteDebit(ctx, logger, testDebitID)
	if err != nil {
		logger.WithError(err).Error("Failed to delete debit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":   "succeeded",
			"debit": deletedDebit,
		})
	}
}
