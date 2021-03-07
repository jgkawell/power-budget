package handler

import (
	"context"
	"time"

	m "api/model"
	s "api/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Handler is the struct for managing all incoming client requests
type Handler struct {
	logger  *logrus.Entry
	router  *gin.Engine
	service s.MetaService
}

// CreateHandler returns a new gin rest handler
func CreateHandler(ctx context.Context, logger *logrus.Entry, config m.Config, service s.MetaService) Handler {

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
		logger:  logger,
		router:  router,
		service: service,
	}

	// Setup routes for accounts
	accounts := h.router.Group("api/accounts")
	{
		accounts.GET("/create", h.CreateAccount)
		accounts.GET("/read", h.ReadAccount)
		accounts.GET("/update", h.UpdateAccount)
		accounts.GET("/delete", h.DeleteAccount)
	}

	// Setup routes for credits
	credits := h.router.Group("api/credits")
	{
		credits.GET("/credits/create", h.CreateCredit)
		credits.GET("/credits/read", h.ReadCredit)
		credits.GET("/credits/update", h.UpdateCredit)
		credits.GET("/credits/delete", h.DeleteCredit)
	}

	// Setup routes for debits
	debits := h.router.Group("api/debits")
	{
		debits.GET("/create", h.CreateDebit)
		debits.GET("/read", h.ReadDebit)
		debits.GET("/update", h.UpdateDebit)
		debits.GET("/delete", h.DeleteDebit)
	}

	return h
}

// RunHandler starts the handler server
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
