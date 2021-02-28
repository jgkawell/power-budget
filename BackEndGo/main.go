package main

import (
	"context"

	"backend/dao"
	"backend/handler"
	"backend/tools"
)

func main() {
	// Create logger and config
	logger, config := tools.GetConfig()

	// Set context for initialization
	ctx := context.Background()

	logger.Info("starting")

	// Create database
	db := dao.CreateConnection(logger, config.DatabaseConfig)
	defer db.Close()

	// Create handlers
	h := handler.CreateRestHandler(ctx, logger, config, db)
	h.RunHandler()
}
