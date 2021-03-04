package dao

import (
	"context"
	"fmt"
	"time"

	m "api/model"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const noResultErrorMsg = "Query operation returned no result. Was the ID not in the database?"

// MetaDao is the wrapper dao around the sub-daos for specific tables
type MetaDao interface {
	Accounts() AccountsDao
	Close()
}

type metaDao struct {
	db       *sqlx.DB
	accounts AccountsDao
}

// CreateDao returns a sqlx db that uses a pgx connection pool
func CreateDao(logger *logrus.Entry, config m.DatabaseConfig) MetaDao {
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
	db := sqlx.NewDb(nativeDB, "pgx")

	conn := metaDao{
		db:       db,
		accounts: NewAccountsDao(db),
	}

	return conn
}

// Accounts returns the accounts sub-dao of the meta dao
func (d metaDao) Accounts() AccountsDao {
	return d.accounts
}

// Close closes the database connection pool
func (d metaDao) Close() {
	d.db.Close()
}

// genericNamedQuery a named query in the database using given SQL and entry
// Params:
//
// - db *sqlx.DB = the db connection to execute the query against
// - sql string = the (formatted) sql statement to execute
// - entry interface{} = the entry with fields to substitute into the sql statement
// - desiredType interface{} = a struct with the (real) type that the DB result will be converted to
//
// Returns:
//
// - interface{} = the result of the query but needs to be converted to specific type (desiredType)
// - error = not nil error if occured
//
func genericNamedQuery(ctx context.Context, logger *logrus.Entry, db *sqlx.DB, sql string, entry interface{}, desiredType interface{}) (interface{}, error) {
	// Run query
	rows, err := db.NamedQuery(sql, entry)
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
	case m.Account:
		logger.WithField("type", t).Debug("Converting to account")
		var convertedResult m.Account
		err = rows.StructScan(&convertedResult)
		return convertedResult, err
	case m.Config:
		// TODO: This is temporary, will be replaced with real DB models
		logger.WithField("type", t).Debug("Converting to config")
		var convertedResult m.Config
		err = rows.StructScan(&convertedResult)
		return convertedResult, err
	case m.DatabaseConfig:
		// TODO: This is temporary, will be replaced with real DB models
		logger.WithField("type", t).Debug("Converting to database config")
		var convertedResult m.DatabaseConfig
		err = rows.StructScan(&convertedResult)
		return convertedResult, err
	default:
		err := fmt.Errorf("unsupported type")
		logger.WithField("type", t).WithError(err).Error("unsupported type")
		return desiredType, err
	}
}
