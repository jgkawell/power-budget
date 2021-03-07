package dao

import (
	"context"

	m "api/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const createDebit = `
	INSERT INTO debits(
		id,
		posted_date,
		amount,
		vendor,
		purpose,
		debit,
		budget,
		notes)
	VALUES(
		:id,
		:posted_date,
		:amount,
		:vendor,
		:purpose,
		:debit,
		:budget,
		:notes)
	RETURNING *;`

const readDebitByID = `
	SELECT *
	FROM debits
	WHERE id = :id;`

const updateDebit = `
	UPDATE debits
	SET posted_date = :posted_date,
		amount = :amount,
		vendor = :vendor,
		purpose = :purpose,
		debit = :debit,
		budget = :budget,
		notes = :notes
	WHERE id = :id
	RETURNING *;`

const deleteDebitByID = `
	DELETE FROM debits
	WHERE id = :id
	RETURNING *;`

// DebitsDao is an interface that enables CRUD operations for acccounts
type DebitsDao interface {
	CreateDebit(ctx context.Context, logger *logrus.Entry, debit m.Debit) (m.Debit, error)
	ReadDebit(ctx context.Context, logger *logrus.Entry, id string) (m.Debit, error)
	UpdateDebit(ctx context.Context, logger *logrus.Entry, debit m.Debit) (m.Debit, error)
	DeleteDebit(ctx context.Context, logger *logrus.Entry, id string) (m.Debit, error)
}

type debitsDao struct {
	db *sqlx.DB
}

// NewDebitsDao returns a new debitsDao struct with given pool connection
func NewDebitsDao(db *sqlx.DB) DebitsDao {
	return debitsDao{db}
}

// CreateDebit creates a debit in the database
func (a debitsDao) CreateDebit(ctx context.Context, logger *logrus.Entry, debit m.Debit) (m.Debit, error) {
	// Set the ID as a new UUID
	debit.ID = uuid.New().String()
	logger = logger.WithField("debit_id", debit.ID)

	// Attempt generic create
	var desiredType m.Debit
	result, err := genericNamedQuery(ctx, logger, a.db, createDebit, debit, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to create debit")
		return m.Debit{}, err
	}

	// Cast a return
	logger.Info("Created debit")
	return result.(m.Debit), nil
}

// ReadDebit reads a debit by id
func (a debitsDao) ReadDebit(ctx context.Context, logger *logrus.Entry, id string) (m.Debit, error) {
	logger = logger.WithField("debit_id", id)

	// Create debit with id to read
	debit := m.Debit{
		ID: id,
	}

	var desiredType m.Debit
	result, err := genericNamedQuery(ctx, logger, a.db, readDebitByID, debit, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to query row")
		return m.Debit{}, err
	}

	// Cast and return
	logger.Info("Read debit")
	return result.(m.Debit), nil
}

// UpdateDebit updates a debit by id with values provided in the struct
// TODO: What happens if all values are not provided? (e.g. Are TEXT fields set to "" in the DB?)
func (a debitsDao) UpdateDebit(ctx context.Context, logger *logrus.Entry, debit m.Debit) (m.Debit, error) {
	logger = logger.WithField("debit_id", debit.ID)

	// Attempt generic update
	var desiredType m.Debit
	result, err := genericNamedQuery(ctx, logger, a.db, updateDebit, debit, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to update debit")
		return m.Debit{}, err
	}

	// Cast and return
	logger.Info("Updated debit")
	return result.(m.Debit), nil
}

// DeleteDebit deletes a debit by id
func (a debitsDao) DeleteDebit(ctx context.Context, logger *logrus.Entry, id string) (m.Debit, error) {
	logger = logger.WithField("debit_id", id)

	// Create debit with id to delete
	debit := m.Debit{
		ID: id,
	}

	var desiredType m.Debit
	result, err := genericNamedQuery(ctx, logger, a.db, deleteDebitByID, debit, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to query row")
		return m.Debit{}, err
	}

	// Cast and return
	logger.Info("Deleted debit")
	return result.(m.Debit), nil
}
