package dao

import (
	"context"
	"time"

	"api/model"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DatabaseConnection interface {
	// CRUD for accounts
	CreateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) model.Account
	ReadAccount(ctx context.Context, logger *logrus.Entry, id uint16) model.Account
	UpdateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) model.Account
	DeleteAccount(ctx context.Context, logger *logrus.Entry, id uint16) bool

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
		logger.WithError(err).Panic("Call to pgx.NewConnPool failed")
	}

	// Then set up sqlx and return the created DB reference
	nativeDB := stdlib.OpenDBFromPool(connPool)

	conn := connection{
		db: sqlx.NewDb(nativeDB, "pgx"),
	}

	return conn
}

// Close closses the database connection pool
func (conn connection) Close() {
	conn.db.Close()
}
