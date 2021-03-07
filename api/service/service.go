package service

import (
	d "api/dao"
)

// MetaService is the wrapper service around the sub-services for specific tables
type MetaService interface {
	Accounts() AccountsService
}

type metaService struct {
	accounts AccountsService
}

// CreateService creates the general service wrapper
func CreateService(metaDao d.MetaDao) MetaService {
	return metaService{
		accounts: CreateAccountsService(metaDao),
	}
}

// Accounts returns the accounts field of the service
func (s metaService) Accounts() AccountsService {
	return s.accounts
}
