package main

import (
	"context"

	d "api/dao"
	h "api/handler"
	s "api/service"
	t "api/tools"
)

func main() {
	// Create logger and config
	logger, config := t.GetConfig()

	// Set context for initialization
	ctx := context.Background()

	logger.Info("starting")

	// Create database
	dao := d.CreateDao(logger, config.DatabaseConfig)
	defer dao.Close()

	// Create service layers
	service := s.CreateService(dao)

	// Create handlers
	handler := h.CreateHandler(ctx, logger, config, service)

	// Start app
	handler.RunHandler()
}
