package handler

import (
	"context"

	"backend/dao"
	"backend/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	router *gin.Engine
	db     dao.DatabaseConnection
}

// CreateRestHandler returns a new gin rest handler
func CreateRestHandler(ctx context.Context, logger *logrus.Entry, config model.Config, db dao.DatabaseConnection) Handler {
	r := gin.Default()
	r.GET("/create", func(c *gin.Context) {
		// Test create dao
		newAccount := model.Account{
			Name:          "Chase Debit 4",
			Balance:       4392.34,
			TotalIn:       34.55,
			TotalOut:      233.44,
			Type:          "Cash",
			CardNumber:    "****",
			AccountNumber: "****-1234",
		}
		createdAccount := db.CreateAccount(ctx, logger, newAccount)
		logger.Info(createdAccount)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/all", func(c *gin.Context) {
		// Test query dao
		allAccounts := db.GetAll(ctx, logger, "accounts")
		logger.Info(allAccounts)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return Handler{r, db}
}

func (h Handler) RunHandler() {
	h.router.Run()
}
