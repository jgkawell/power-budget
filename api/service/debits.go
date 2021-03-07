package service

import (
	"context"

	d "api/dao"
	m "api/model"

	"github.com/sirupsen/logrus"
)

// DebitsService is the interface for managing the debits table
type DebitsService interface {
	CreateDebit(ctx context.Context, logger *logrus.Entry, account m.Debit) (m.Debit, error)
	ReadDebit(ctx context.Context, logger *logrus.Entry, id string) (m.Debit, error)
	UpdateDebit(ctx context.Context, logger *logrus.Entry, account m.Debit) (m.Debit, error)
	DeleteDebit(ctx context.Context, logger *logrus.Entry, id string) (m.Debit, error)
}

type debitsService struct {
	dao d.MetaDao
}

// CreateDebitsService creates the debits service
func CreateDebitsService(metaDao d.MetaDao) DebitsService {
	return debitsService{metaDao}
}

func (s debitsService) CreateDebit(ctx context.Context, logger *logrus.Entry, account m.Debit) (m.Debit, error) {
	return s.dao.Debits().CreateDebit(ctx, logger, account)
}

func (s debitsService) ReadDebit(ctx context.Context, logger *logrus.Entry, id string) (m.Debit, error) {
	return s.dao.Debits().ReadDebit(ctx, logger, id)
}

func (s debitsService) UpdateDebit(ctx context.Context, logger *logrus.Entry, account m.Debit) (m.Debit, error) {
	return s.dao.Debits().UpdateDebit(ctx, logger, account)
}

func (s debitsService) DeleteDebit(ctx context.Context, logger *logrus.Entry, id string) (m.Debit, error) {
	return s.dao.Debits().DeleteDebit(ctx, logger, id)
}
