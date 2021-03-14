package service

import (
	"context"

	d "api/dao"
	m "api/model"

	"github.com/sirupsen/logrus"
)

// AccountsService is the interface for managing the accounts table
type AccountsService interface {
	CreateAccount(ctx context.Context, logger *logrus.Entry, account m.Account) (m.Account, error)
	ReadAccount(ctx context.Context, logger *logrus.Entry, id string) (m.Account, error)
	ReadAllAccounts(ctx context.Context, logger *logrus.Entry) ([]m.Account, error)
	UpdateAccount(ctx context.Context, logger *logrus.Entry, account m.Account) (m.Account, error)
	DeleteAccount(ctx context.Context, logger *logrus.Entry, id string) (m.Account, error)
}

type accountsService struct {
	dao d.MetaDao
}

// CreateAccountsService creates the accounts service
func CreateAccountsService(metaDao d.MetaDao) AccountsService {
	return accountsService{metaDao}
}

func (s accountsService) CreateAccount(ctx context.Context, logger *logrus.Entry, account m.Account) (m.Account, error) {
	return s.dao.Accounts().CreateAccount(ctx, logger, account)
}

func (s accountsService) ReadAccount(ctx context.Context, logger *logrus.Entry, id string) (m.Account, error) {
	return s.dao.Accounts().ReadAccount(ctx, logger, id)
}

func (s accountsService) ReadAllAccounts(ctx context.Context, logger *logrus.Entry) ([]m.Account, error) {
	return s.dao.Accounts().ReadAllAccounts(ctx, logger)
}

func (s accountsService) UpdateAccount(ctx context.Context, logger *logrus.Entry, account m.Account) (m.Account, error) {
	return s.dao.Accounts().UpdateAccount(ctx, logger, account)
}

func (s accountsService) DeleteAccount(ctx context.Context, logger *logrus.Entry, id string) (m.Account, error) {
	return s.dao.Accounts().DeleteAccount(ctx, logger, id)
}
