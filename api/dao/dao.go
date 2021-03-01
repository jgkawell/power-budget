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

const noResultErrorMsg = "Query operation returned no result"

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

// Executes a named query in the database using given SQL and object
func (conn connection) NamedQuery(ctx context.Context, logger *logrus.Entry, sql string, object interface{}) (*sqlx.Rows, error) {
	// Run query
	rows, err := conn.db.NamedQuery(sql, object)
	if err != nil {
		logger.WithError(err).Error("Error executing named query")
		return nil, err
	}

	// Attempt to get next result
	next := rows.Next()
	if !next {
		logger.Error(noResultErrorMsg)
		return nil, fmt.Errorf(noResultErrorMsg)
	}

	// Return result
	logger.Debug("Succeeded executing named query")
	return rows, nil
}
