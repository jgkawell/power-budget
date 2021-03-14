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
	credit, err := h.service.Credits().CreateCredit(ctx, logger, newCredit)
	if err != nil {
		logger.WithError(err).Error("Failed to create credit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, credit)
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

	// Call dao layer
	credit, err := h.service.Credits().ReadCredit(ctx, logger, id)
	if err != nil {
		logger.WithError(err).Error("Failed to read credit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, credit)
	}
}

// ReadAllCredits reads all debits from the database
func (h Handler) ReadAllCredits(ctx *gin.Context) {
	logger := h.logger.WithField("handler", "ReadAllCredits")

	// Call dao layer
	debits, err := h.service.Credits().ReadAllCredits(ctx, logger)
	if err != nil {
		logger.WithError(err).Error("Failed to read all credits")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, debits)
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
	credit, err := h.service.Credits().UpdateCredit(ctx, logger, newCredit)
	if err != nil {
		logger.WithError(err).Error("Failed to update credit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, credit)
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

	// Call dao layer
	credit, err := h.service.Credits().DeleteCredit(ctx, logger, id)
	if err != nil {
		logger.WithError(err).Error("Failed to delete credit")
		ctx.JSON(500, gin.H{
			"msg": "failed",
		})
	} else {
		ctx.JSON(200, credit)
	}
}
