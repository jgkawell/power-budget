package service

import (
	"context"

	d "api/dao"
	m "api/model"

	"github.com/sirupsen/logrus"
)

// CreditsService is the interface for managing the credits table
type CreditsService interface {
	CreateCredit(ctx context.Context, logger *logrus.Entry, account m.Credit) (m.Credit, error)
	ReadCredit(ctx context.Context, logger *logrus.Entry, id string) (m.Credit, error)
	UpdateCredit(ctx context.Context, logger *logrus.Entry, account m.Credit) (m.Credit, error)
	DeleteCredit(ctx context.Context, logger *logrus.Entry, id string) (m.Credit, error)
}

type creditsService struct {
	dao d.MetaDao
}

// CreateCreditsService creates the credits service
func CreateCreditsService(metaDao d.MetaDao) CreditsService {
	return creditsService{metaDao}
}

func (s creditsService) CreateCredit(ctx context.Context, logger *logrus.Entry, account m.Credit) (m.Credit, error) {
	return s.dao.Credits().CreateCredit(ctx, logger, account)
}

func (s creditsService) ReadCredit(ctx context.Context, logger *logrus.Entry, id string) (m.Credit, error) {
	return s.dao.Credits().ReadCredit(ctx, logger, id)
}

func (s creditsService) UpdateCredit(ctx context.Context, logger *logrus.Entry, account m.Credit) (m.Credit, error) {
	return s.dao.Credits().UpdateCredit(ctx, logger, account)
}

func (s creditsService) DeleteCredit(ctx context.Context, logger *logrus.Entry, id string) (m.Credit, error) {
	return s.dao.Credits().DeleteCredit(ctx, logger, id)
}
