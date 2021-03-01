package dao

import (
	"context"

	"api/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const create = `
	INSERT INTO accounts(
		id,
		name,
		balance,
		total_in,
		total_out,
		type,
		card_number,
		account_number)
	VALUES(
		:id,
		:name,
		:balance,
		:total_in,
		:total_out,
		:type,
		:card_number,
		:account_number)
	RETURNING *;`

const read = `
	SELECT *
	FROM accounts
	WHERE id = $1;`

const update = `
	UPDATE accounts
	SET name = :name,
		balance = :balance,
		total_in = :total_in,
		total_out = :total_out,
		type = :type,
		card_number = :card_number,
		account_number = :account_number
	WHERE id = :id;`

const delete = `
	DELETE FROM accounts
	WHERE id = $1;`

// CreateAccount creates an account in the database and returns the id if succeeded
func (conn connection) CreateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) (model.Account, error) {

	// Set the ID as a new UUID
	account.ID = uuid.New().String()

	// Attempt generic create
	rows, err := conn.NamedQuery(ctx, logger, create, account)
	if err != nil {
		logger.WithError(err).Error("Failed to create account")
		return model.Account{}, err
	}

	return buildAccountFromRows(logger, rows)
}

// ReadAccount reads an account by id
func (conn connection) ReadAccount(ctx context.Context, logger *logrus.Entry, id uint16) model.Account {

	// Run read query
	var readAccount model.Account
	err := conn.db.Get(&readAccount, read, id)
	if err != nil {
		logger.WithError(err).Error("Failed to query row")
		return model.Account{}
	}

	// Return result
	logger.Info("Read succeeded")
	return readAccount
}

// ReadAccount reads an account by id
func (conn connection) UpdateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) (model.Account, error) {

	// Attempt generic update
	rows, err := conn.NamedQuery(ctx, logger, update, account)
	if err != nil {
		logger.WithError(err).Error("Failed to update account")
		return model.Account{}, err
	}

	return buildAccountFromRows(logger, rows)
}

// DeleteAccount deletes an account by id
func (conn connection) DeleteAccount(ctx context.Context, logger *logrus.Entry, id string) bool {
	logger = logger.WithField("account_id", id)

	// Run delete query
	rows, err := conn.db.Exec(delete, id)
	if err == nil {
		count, err := rows.RowsAffected()
		if err == nil {
			countLogger := logger.WithField("count", count)
			if count == 1 {
				countLogger.Info("Delete succeeded")
				return true
			} else if count > 1 {
				countLogger.Warn("Deleted multiple records with single ID. Records should not have duplicate IDs.")
				return true
			} else if count == 0 {
				countLogger.Debug("Nothing was deleted. Was the ID not in the DB?")
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

func buildAccountFromRows(logger *logrus.Entry, rows *sqlx.Rows) (model.Account, error) {
	// Read out row into account struct
	var account model.Account
	err := rows.StructScan(&account)
	if err != nil {
		logger.WithError(err).Error("Failed to scan struct from rows")
		return model.Account{}, err
	}
	// Return result
	return account, nil
}
