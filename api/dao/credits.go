package dao

import (
	"context"

	m "api/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const createCredit = `
	INSERT INTO credits(
		id,
		posted_date,
		amount,
		source,
		purpose,
		credit,
		budget,
		notes)
	VALUES(
		:id,
		:posted_date,
		:amount,
		:source,
		:purpose,
		:credit,
		:budget,
		:notes)
	RETURNING *;`

const readCreditByID = `
	SELECT *
	FROM credits
	WHERE id = :id;`

const updateCredit = `
	UPDATE credits
	SET posted_date = :posted_date,
		amount = :amount,
		source = :source,
		purpose = :purpose,
		credit = :credit,
		budget = :budget,
		notes = :notes
	WHERE id = :id
	RETURNING *;`

const deleteCreditByID = `
	DELETE FROM credits
	WHERE id = :id
	RETURNING *;`

// CreditsDao is an interface that enables CRUD operations for acccounts
type CreditsDao interface {
	CreateCredit(ctx context.Context, logger *logrus.Entry, credit m.Credit) (m.Credit, error)
	ReadCredit(ctx context.Context, logger *logrus.Entry, id string) (m.Credit, error)
	UpdateCredit(ctx context.Context, logger *logrus.Entry, credit m.Credit) (m.Credit, error)
	DeleteCredit(ctx context.Context, logger *logrus.Entry, id string) (m.Credit, error)
}

type creditsDao struct {
	db *sqlx.DB
}

// NewCreditsDao returns a new creditsDao struct with given pool connection
func NewCreditsDao(db *sqlx.DB) CreditsDao {
	return creditsDao{db}
}

// CreateCredit creates a credit in the database
func (a creditsDao) CreateCredit(ctx context.Context, logger *logrus.Entry, credit m.Credit) (m.Credit, error) {
	// Set the ID as a new UUID
	credit.ID = uuid.New().String()
	logger = logger.WithField("credit_id", credit.ID)

	// Attempt generic create
	var desiredType m.Credit
	result, err := genericNamedQuery(ctx, logger, a.db, createCredit, credit, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to create credit")
		return m.Credit{}, err
	}

	// Cast a return
	logger.Info("Created credit")
	return result.(m.Credit), nil
}

// ReadCredit reads a credit by id
func (a creditsDao) ReadCredit(ctx context.Context, logger *logrus.Entry, id string) (m.Credit, error) {
	logger = logger.WithField("credit_id", id)

	// Create credit with id to read
	credit := m.Credit{
		ID: id,
	}

	var desiredType m.Credit
	result, err := genericNamedQuery(ctx, logger, a.db, readCreditByID, credit, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to query row")
		return m.Credit{}, err
	}

	// Cast and return
	logger.Info("Read credit")
	return result.(m.Credit), nil
}

// UpdateCredit updates a credit by id with values provided in the struct
// TODO: What happens if all values are not provided? (e.g. Are TEXT fields set to "" in the DB?)
func (a creditsDao) UpdateCredit(ctx context.Context, logger *logrus.Entry, credit m.Credit) (m.Credit, error) {
	logger = logger.WithField("credit_id", credit.ID)

	// Attempt generic update
	var desiredType m.Credit
	result, err := genericNamedQuery(ctx, logger, a.db, updateCredit, credit, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to update credit")
		return m.Credit{}, err
	}

	// Cast and return
	logger.Info("Updated credit")
	return result.(m.Credit), nil
}

// DeleteCredit deletes a credit by id
func (a creditsDao) DeleteCredit(ctx context.Context, logger *logrus.Entry, id string) (m.Credit, error) {
	logger = logger.WithField("credit_id", id)

	// Create credit with id to delete
	credit := m.Credit{
		ID: id,
	}

	var desiredType m.Credit
	result, err := genericNamedQuery(ctx, logger, a.db, deleteCreditByID, credit, desiredType)
	if err != nil {
		logger.WithError(err).Error("Failed to query row")
		return m.Credit{}, err
	}

	// Cast and return
	logger.Info("Deleted credit")
	return result.(m.Credit), nil
}
