package service

import (
	d "api/dao"
)

// MetaService is the wrapper service around the sub-services for specific tables
type MetaService interface {
	Accounts() AccountsService
	Credits() CreditsService
	Debits() DebitsService
}

type metaService struct {
	accounts AccountsService
	credits  CreditsService
	debits   DebitsService
}

// CreateService creates the general service wrapper
func CreateService(metaDao d.MetaDao) MetaService {
	return metaService{
		accounts: CreateAccountsService(metaDao),
		credits:  CreateCreditsService(metaDao),
		debits:   CreateDebitsService(metaDao),
	}
}

// Accounts returns the accounts field of the service
func (s metaService) Accounts() AccountsService {
	return s.accounts
}

// Credits returns the credits field of the service
func (s metaService) Credits() CreditsService {
	return s.credits
}

// Debits returns the debits field of the service
func (s metaService) Debits() DebitsService {
	return s.debits
}
