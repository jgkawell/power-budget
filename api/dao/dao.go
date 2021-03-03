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
	ReadAccount(ctx context.Context, logger *logrus.Entry, id string) (model.Account, error)
	UpdateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) (model.Account, error)
	DeleteAccount(ctx context.Context, logger *logrus.Entry, id string) (model.Account, error)

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
//
// - sql string = the (formatted) sql statement to execute
// - entry interface{} = the entry with fields to substitute into the sql statement
// - desiredType interface{} = a struct with the (real) type that the DB result will be converted to
//
// Returns:
// - interface{} = the result of the query but needs to be converted to specific type (desiredType)
// - error = not nil error if occured
func (conn connection) genericNamedQuery(ctx context.Context, logger *logrus.Entry, sql string, entry interface{}, desiredType interface{}) (interface{}, error) {
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
	switch t := desiredType.(type) {
	case model.Account:
		logger.WithField("type", t).Debug("Converting to account")
		var convertedResult model.Account
		err = rows.StructScan(&convertedResult)
		return convertedResult, err
	case model.Config:
		// TODO: This is temporary, will be replaced with real DB models
		logger.WithField("type", t).Debug("Converting to config")
		var convertedResult model.Config
		err = rows.StructScan(&convertedResult)
		return convertedResult, err
	case model.DatabaseConfig:
		// TODO: This is temporary, will be replaced with real DB models
		logger.WithField("type", t).Debug("Converting to database config")
		var convertedResult model.DatabaseConfig
		err = rows.StructScan(&convertedResult)
		return convertedResult, err
	default:
		err := fmt.Errorf("unsupported type")
		logger.WithField("type", t).WithError(err).Error("unsupported type")
		return desiredType, err
	}
}
