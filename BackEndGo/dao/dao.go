package dao

import (
	"context"
	"fmt"
	"time"

	"backend/model"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const insert = `INSERT INTO accounts(name, type, card_number, account_number, balance, total_in, total_out)
	VALUES(:name, :type, :card_number, :account_number, :balance, :total_in, :total_out) RETURNING *;`

type DatabaseConnection interface {
	GetAll(ctx context.Context, logger *logrus.Entry, table string) []model.Account
	CreateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) model.Account
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

// GetAll returns all rows for a given table
func (conn connection) GetAll(ctx context.Context, logger *logrus.Entry, table string) []model.Account {
	accounts := []model.Account{}
	err := conn.db.Select(&accounts, fmt.Sprintf("select * from %s", table))
	if err != nil {
		logger.WithError(err).Error("Failed to query")
	}
	return accounts
}

// CreateAccount creates an account in the database and returns the result
func (conn connection) CreateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) model.Account {

	// Run named query insert
	rows, err := conn.db.NamedQuery(insert, account)
	if err != nil {
		logger.WithError(err).Error("Failed to query row")
		return model.Account{}
	}

	// Attempt to get next result
	next := rows.Next()
	if !next {
		logger.Error("Create returned no result")
		return model.Account{}
	}

	// Read out row into account struct
	var createdAccount model.Account
	err = rows.StructScan(&createdAccount)
	if err != nil {
		logger.WithError(err).Error("Failed to scan struct from rows")
	}

	// Return result
	logger.WithField("result", createdAccount).Info("Insert succeeded")
	return createdAccount
}

// Close closses the database connection pool
func (conn connection) Close() {
	conn.db.Close()
}
