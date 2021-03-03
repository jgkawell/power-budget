package dao

import (
	"context"

	"api/model"

	"github.com/google/uuid"
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
	var result model.Account
	newAccount, err := conn.genericNamedQuery(ctx, logger, create, account, result)
	if err != nil {
		logger.WithError(err).Error("Failed to create account")
		return model.Account{}, err
	}

	// Cast a return
	result = newAccount.(model.Account)
	return result, nil
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
	var result model.Account
	updatedAccount, err := conn.genericNamedQuery(ctx, logger, update, account, result)
	if err != nil {
		logger.WithError(err).Error("Failed to update account")
		return model.Account{}, err
	}

	// Cast and return
	result = updatedAccount.(model.Account)
	return result, nil
}

// DeleteAccount deletes an account by id
func (conn connection) DeleteAccount(ctx context.Context, logger *logrus.Entry, id string) bool {
	logger = logger.WithField("account_id", id)

	// Run delete query
	return conn.genericDelete(ctx, logger, delete, id)
}
