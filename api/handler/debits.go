package handler

import (
	"net/http"

	m "api/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateDebit creates a debit in the database
func (h Handler) CreateDebit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "CreateDebit")

	// Pull out debit from request
	var newDebit m.Debit
	if err := ctx.ShouldBindJSON(&newDebit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call dao layer
	debit, err := h.service.Debits().CreateDebit(ctx, logger, newDebit)
	if err != nil {
		logger.WithError(err).Error("Failed to create debit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, debit)
	}
}

// ReadDebit reads a debit from the database with given id
func (h Handler) ReadDebit(ctx *gin.Context) {
	// Pull out id to read and add to logger
	id := ctx.Param("id")
	logger := h.logger.WithFields(logrus.Fields{
		"debit_id": id,
		"handler":  "ReadDebit",
	})

	// Call dao layer
	debit, err := h.service.Debits().ReadDebit(ctx, logger, id)
	if err != nil {
		logger.WithError(err).Error("Failed to read debit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, debit)
	}
}

// ReadAllDebits reads all debits from the database
func (h Handler) ReadAllDebits(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "ReadAllDebits")

	// Call dao layer
	debits, err := h.service.Debits().ReadAllDebits(ctx, logger)
	if err != nil {
		logger.WithError(err).Error("Failed to read all debits")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, debits)
	}
}

// UpdateDebit updates a debit in the database with given parameters (id required)
func (h Handler) UpdateDebit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "UpdateDebit")

	// Pull out debit from request
	var newDebit m.Debit
	if err := ctx.ShouldBindJSON(&newDebit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger = logger.WithField("debit_id", newDebit.ID)

	// Call dao layer
	debit, err := h.service.Debits().UpdateDebit(ctx, logger, newDebit)
	if err != nil {
		logger.WithError(err).Error("Failed to update debit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, debit)
	}
}

// DeleteDebit deletes a debit in the database with given id
func (h Handler) DeleteDebit(ctx *gin.Context) {
	// Pull out id to delete and add to logger
	id := ctx.Param("id")
	logger := h.logger.WithFields(logrus.Fields{
		"debit_id": id,
		"handler":  "DeleteDebit",
	})

	// Call dao layer
	debit, err := h.service.Debits().DeleteDebit(ctx, logger, id)
	if err != nil {
		logger.WithError(err).Error("Failed to delete debit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, debit)
	}
}
