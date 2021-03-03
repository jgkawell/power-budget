package dao

import (
	"context"
	"fmt"
	"time"

	"api/model"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const noResultErrorMsg = "Query operation returned no result. Was the ID not in the database?"

type DatabaseConnection interface {
	// CRUD for accounts
	CreateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) (model.Account, error)
	ReadAccount(ctx context.Context, logger *logrus.Entry, id uint16) model.Account
	UpdateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) (model.Account, error)
	DeleteAccount(ctx context.Context, logger *logrus.Entry, id string) bool

	// Helpers
	Close()
}

type connection struct {
	db *sqlx.DB
}

// CreateConnection returns a sqlx db that uses a pgx connection pool
func CreateConnection(logger *logrus.Entry, config model.DatabaseConfig) DatabaseConnection {
	// First set up the pgx connection pool
	connConfig := pgx.ConnConfig{
		Host:     config.Host,
		Port:     config.Port,
		Database: config.Database,
		User:     config.User,
		Password: config.Password,
	}
	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		AfterConnect:   nil,
		MaxConnections: 20,
		AcquireTimeout: 30 * time.Second,
	})
	if err != nil {
		logger.WithError(err).Fatal("Call to pgx.NewConnPool failed")
	}

	// Then set up sqlx and return the created DB reference
	nativeDB := stdlib.OpenDBFromPool(connPool)

	conn := connection{
		db: sqlx.NewDb(nativeDB, "pgx"),
	}

	return conn
}

// Close closes the database connection pool
func (conn connection) Close() {
	conn.db.Close()
}

// Executes a named query in the database using given SQL and entry
// Params:
// - sql = the (formatted) sql statement to execute
// - entry = the entry with fields to substitute into the sql statement
// Returns:
// - interface{} = the result of the query but needs to be converted back to original type
// - error = error if occured
func (conn connection) genericNamedQuery(ctx context.Context, logger *logrus.Entry, sql string, entry interface{}, result interface{}) (interface{}, error) {
	// Run query
	rows, err := conn.db.NamedQuery(sql, entry)
	if err != nil {
		logger.WithError(err).Error("Error executing named query")
		return nil, err
	}

	// Attempt to get next result
	next := rows.Next()
	if !next {
		err := fmt.Errorf(noResultErrorMsg)
		logger.WithError(err).Error(noResultErrorMsg)
		return nil, err
	}

	// Convert to model struct from interface
	switch v := result.(type) {
	case model.Account:
		logger.WithField("type", v).Debug("Converting to account")
		var convertedResult model.Account
		err = rows.StructScan(&convertedResult)
		return convertedResult, err
	case model.Config:
		logger.WithField("type", v).Debug("Converting to config")
		var convertedResult model.Config
		err = rows.StructScan(&convertedResult)
		return convertedResult, err
	case model.DatabaseConfig:
		logger.WithField("type", v).Debug("Converting to database config")
		var convertedResult model.DatabaseConfig
		err = rows.StructScan(&convertedResult)
		return convertedResult, err
	default:
		err := fmt.Errorf("unsupported type")
		logger.WithField("type", v).WithError(err).Error("unsupported type")
		return result, err
	}
}

// genericDelete deletes any record from the database with given sql and id
func (conn connection) genericDelete(ctx context.Context, logger *logrus.Entry, sql string, id string) bool {
	// Run delete query
	result, err := conn.db.Exec(sql, id)
	if err == nil {
		count, err := result.RowsAffected()
		if err == nil {
			countLogger := logger.WithField("count", count)
			if count == 1 {
				countLogger.Info("Delete succeeded")
				return true
			} else if count > 1 {
				countLogger.Warn("Deleted multiple records with single ID. Records should not have duplicate IDs.")
				return true
			} else if count == 0 {
				countLogger.Warn("Nothing was deleted. Was the ID not in the DB?")
				return false
			} else {
				countLogger.Error("Look at count field. This should never happen.")
				return false
			}
		}
		logger.WithError(err).Error("Failed to get count")
		return false
	}

	// Return result
	logger.WithError(err).Error("Delete failed")
	return false
}
