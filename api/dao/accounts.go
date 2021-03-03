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
	WHERE id = :id;`

const update = `
	UPDATE accounts
	SET name = :name,
		balance = :balance,
		total_in = :total_in,
		total_out = :total_out,
		type = :type,
		card_number = :card_number,
		account_number = :account_number
	WHERE id = :id
	RETURNING *;`

const delete = `
	DELETE FROM accounts
	WHERE id = :id
	RETURNING *;`

// CreateAccount creates an account in the database and returns the id if succeeded
func (conn connection) CreateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) (model.Account, error) {
	// Set the ID as a new UUID
	account.ID = uuid.New().String()
	logger = logger.WithField("account_id", account.ID)

	// Attempt generic create
	var desiredType model.Account
	result, err := conn.genericNamedQuery(ctx, logger, create, account, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to create account")
		return model.Account{}, err
	}

	// Cast a return
	logger.Info("Created account")
	return result.(model.Account), nil
}

// ReadAccount reads an account by id
func (conn connection) ReadAccount(ctx context.Context, logger *logrus.Entry, id string) (model.Account, error) {
	logger = logger.WithField("account_id", id)

	// Create account with id to read
	account := model.Account{
		ID: id,
	}

	var desiredType model.Account
	result, err := conn.genericNamedQuery(ctx, logger, read, account, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to query row")
		return model.Account{}, err
	}

	// Cast and return
	logger.Info("Read account")
	return result.(model.Account), nil
}

// ReadAccount reads an account by id
func (conn connection) UpdateAccount(ctx context.Context, logger *logrus.Entry, account model.Account) (model.Account, error) {
	logger = logger.WithField("account_id", account.ID)

	// Attempt generic update
	var desiredType model.Account
	result, err := conn.genericNamedQuery(ctx, logger, update, account, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to update account")
		return model.Account{}, err
	}

	// Cast and return
	logger.Info("Updated account")
	return result.(model.Account), nil
}

// DeleteAccount deletes an account by id
func (conn connection) DeleteAccount(ctx context.Context, logger *logrus.Entry, id string) (model.Account, error) {
	logger = logger.WithField("account_id", id)

	// Create account with id to delete
	account := model.Account{
		ID: id,
	}

	var desiredType model.Account
	result, err := conn.genericNamedQuery(ctx, logger, delete, account, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to query row")
		return model.Account{}, err
	}

	// Cast and return
	logger.Info("Deleted account")
	return result.(model.Account), nil
}
