package handler

import (
	"net/http"

	m "api/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateCredit creates a credit in the database
func (h Handler) CreateCredit(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "CreateCredit")

	// Pull out credit from request
	var newCredit m.Credit
	if err := ctx.ShouldBindJSON(&newCredit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call dao layer
	createdCredit, err := h.service.Credits().CreateCredit(ctx, logger, newCredit)
	if err != nil {
		logger.WithError(err).Error("Failed to create credit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":    "succeeded",
			"credit": createdCredit,
		})
	}
}

// ReadCredit reads a credit from the database with given id
func (h Handler) ReadCredit(ctx *gin.Context) {
	// Pull out id to read and add to logger
	id := ctx.Param("id")
	logger := h.logger.WithFields(logrus.Fields{
		"credit_id": id,
		"handler":   "ReadCredit",
	})
	logger.Debug("handler")

	// Call dao layer
	readCredit, err := h.service.Credits().ReadCredit(ctx, logger, id)
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

	// Pull out credit from request
	var newCredit m.Credit
	if err := ctx.ShouldBindJSON(&newCredit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger = logger.WithField("credit_id", newCredit.ID)

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
	// Pull out id to delete and add to logger
	id := ctx.Param("id")
	logger := h.logger.WithFields(logrus.Fields{
		"credit_id": id,
		"handler":   "DeleteCredit",
	})
	logger.Debug("handler")

	// Call dao layer
	deletedCredit, err := h.service.Credits().DeleteCredit(ctx, logger, id)
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
