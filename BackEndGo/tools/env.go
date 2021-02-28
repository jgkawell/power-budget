package tools

import (
	"os"
	"strconv"

	"backend/model"

	"github.com/sirupsen/logrus"
)

const serviceName = "api"

// GetConfig builds the config struct from environment variables
func GetConfig() (*logrus.Entry, model.Config) {

	// Create service wide global logger
	logger := createLogger()

	// Get the database port from the environment
	port, err := strconv.ParseUint(getEnv(logger, "DB_PORT"), 10, 16)
	if err != nil {
		logger.WithError(err).Panic("Failed to convert port to uint")
	}

	// Build config from environment variables
	config := model.Config{
		Env: getEnv(logger, "ENV"),
		DatabaseConfig: model.DatabaseConfig{
			Host:     getEnv(logger, "DB_HOST"),
			Port:     uint16(port),
			Database: getEnv(logger, "DB_DATABASE"),
			User:     getEnv(logger, "DB_USER"),
			Password: getEnv(logger, "DB_PASSWORD"),
		},
	}

	return logger, config
}

// createLogger builds out a logrus global logger using environment variables
func createLogger() *logrus.Entry {
	// Create temp logger to grab environment
	tempLogger := createBasicLogger().WithField("service", serviceName)
	env := getEnv(tempLogger, "ENV")

	// Create global logger and set formatter based on environment
	newLogger := createBasicLogger()
	if env == "dev" {
		newLogger.SetFormatter(&logrus.TextFormatter{})
	}

	// Attempt to read log level: default is INFO
	logLevelString := getEnv(newLogger.WithField("service", serviceName), "LOG_LEVEL")
	logLevel, err := logrus.ParseLevel(logLevelString)
	if err != nil {
		tempLogger.WithError(err).WithFields(logrus.Fields{
			"service":   serviceName,
			"env_level": logLevelString,
		}).Error("Failed to parse log level from env. Defaulting to INFO")
		logLevel = logrus.InfoLevel
	}

	// Create final logger entry and return with service field
	newLogger.SetLevel(logLevel)
	logger := newLogger.WithField("service", serviceName)
	return logger
}

// createBasicLogger builds a simple logrus logger with default values
func createBasicLogger() *logrus.Logger {
	newLogger := logrus.New()
	newLogger.SetFormatter(&logrus.JSONFormatter{})
	// TODO: Find out why this is doing func AND file (only want func)
	newLogger.SetReportCaller(true)
	newLogger.SetOutput(os.Stdout)
	newLogger.SetLevel(logrus.InfoLevel)

	return newLogger
}

// getEnv attempts to retreive the value of an environment variable
// from the OS and panics if not found
func getEnv(logger *logrus.Entry, key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		logger.WithField("key", key).Panic("Required environment variable not found for key")
	}
	return value
}
