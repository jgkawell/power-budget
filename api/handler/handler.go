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

	// Setup routes
	h.router.GET("/account/create", h.CreateAccount)
	h.router.GET("/account/read", h.ReadAccount)
	h.router.GET("/account/update", h.UpdateAccount)
	h.router.GET("/account/delete", h.DeleteAccount)

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
