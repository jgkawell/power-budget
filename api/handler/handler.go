package handler

import (
	"context"
	"time"

	"api/dao"
	"api/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger *logrus.Entry
	db     dao.DatabaseConnection
	router *gin.Engine
}

// CreateRestHandler returns a new gin rest handler
func CreateRestHandler(ctx context.Context, logger *logrus.Entry, config model.Config, db dao.DatabaseConnection) Handler {

	// Create gin router
	if config.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(
		ginrus(logger),
		gin.Recovery(),
	)

	// Create handler
	h := Handler{
		logger: logger,
		db:     db,
		router: router,
	}

	// Setup routes
	h.router.GET("/account/create", h.CreateAccount)
	h.router.GET("/account/read", h.ReadAccount)
	h.router.GET("/account/update", h.UpdateAccount)
	h.router.GET("/account/delete", h.DeleteAccount)

	return h
}

func (h Handler) RunHandler() {
	h.router.Run()
}

// ginrus returns a gin.HandlerFunc (middleware) that logs requests using logrus.
// Credit: https://github.com/gin-gonic/contrib/blob/master/ginrus/ginrus.go
func ginrus(logger *logrus.Entry) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		ctx.Next()
		end := time.Now()
		latency := end.Sub(start)

		entry := logger.WithFields(logrus.Fields{
			"status":     ctx.Writer.Status(),
			"method":     ctx.Request.Method,
			"path":       path,
			"ip":         ctx.ClientIP(),
			"latency":    latency,
			"user-agent": ctx.Request.UserAgent(),
		})

		if len(ctx.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(ctx.Errors.String())
		} else {
			entry.Trace()
		}
	}
}
