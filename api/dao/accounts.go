package dao

import (
	"context"

	m "api/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const createAccount = `
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

const readAccountByID = `
	SELECT *
	FROM accounts
	WHERE id = :id;`

const updateAccount = `
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

const deleteAccountByID = `
	DELETE FROM accounts
	WHERE id = :id
	RETURNING *;`

// AccountsDao is an interface that enables CRUD operations for acccounts
type AccountsDao interface {
	CreateAccount(ctx context.Context, logger *logrus.Entry, account m.Account) (m.Account, error)
	ReadAccount(ctx context.Context, logger *logrus.Entry, id string) (m.Account, error)
	UpdateAccount(ctx context.Context, logger *logrus.Entry, account m.Account) (m.Account, error)
	DeleteAccount(ctx context.Context, logger *logrus.Entry, id string) (m.Account, error)
}

type accountsDao struct {
	db *sqlx.DB
}

// NewAccountsDao returns a new accountsDao struct with given pool connection
func NewAccountsDao(db *sqlx.DB) AccountsDao {
	return accountsDao{db}
}

// CreateAccount creates an account in the database
func (a accountsDao) CreateAccount(ctx context.Context, logger *logrus.Entry, account m.Account) (m.Account, error) {
	// Set the ID as a new UUID
	account.ID = uuid.New().String()
	logger = logger.WithField("account_id", account.ID)

	// Attempt generic create
	var desiredType m.Account
	result, err := genericNamedQuery(ctx, logger, a.db, createAccount, account, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to create account")
		return m.Account{}, err
	}

	// Cast a return
	logger.Info("Created account")
	return result.(m.Account), nil
}

// ReadAccount reads an account by id
func (a accountsDao) ReadAccount(ctx context.Context, logger *logrus.Entry, id string) (m.Account, error) {
	logger = logger.WithField("account_id", id)

	// Create account with id to read
	account := m.Account{
		ID: id,
	}

	var desiredType m.Account
	result, err := genericNamedQuery(ctx, logger, a.db, readAccountByID, account, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to read account")
		return m.Account{}, err
	}

	// Cast and return
	logger.Info("Read account")
	return result.(m.Account), nil
}

// UpdateAccount updates an account by id with values provided in the struct
func (a accountsDao) UpdateAccount(ctx context.Context, logger *logrus.Entry, account m.Account) (m.Account, error) {
	logger = logger.WithField("account_id", account.ID)

	// Attempt generic update
	var desiredType m.Account
	result, err := genericNamedQuery(ctx, logger, a.db, updateAccount, account, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to update account")
		return m.Account{}, err
	}

	// Cast and return
	logger.Info("Updated account")
	return result.(m.Account), nil
}

// DeleteAccount deletes an account by id
func (a accountsDao) DeleteAccount(ctx context.Context, logger *logrus.Entry, id string) (m.Account, error) {
	logger = logger.WithField("account_id", id)

	// Create account with id to delete
	account := m.Account{
		ID: id,
	}

	var desiredType m.Account
	result, err := genericNamedQuery(ctx, logger, a.db, deleteAccountByID, account, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to delete account")
		return m.Account{}, err
	}

	// Cast and return
	logger.Info("Deleted account")
	return result.(m.Account), nil
}
