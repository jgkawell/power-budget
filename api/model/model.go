package model

import (
	"time"
)

// Config contains config values for the api service
type Config struct {
	Env            string
	DatabaseConfig DatabaseConfig
}

// DatabaseConfig holds connection info for PostgreSQL database
type DatabaseConfig struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

// Account defines an account in the database
type Account struct {
	ID            string  `db:"id"`
	Name          string  `db:"name"`
	Balance       float32 `db:"balance"`
	TotalIn       float32 `db:"total_in"`
	TotalOut      float32 `db:"total_out"`
	Type          string  `db:"type"`
	CardNumber    string  `db:"card_number"`
	AccountNumber string  `db:"account_number"`
}

// Credit defines a credit in the database
type Credit struct {
	ID         string    `db:"id"`
	PostedDate time.Time `db:"posted_date"`
	Amount     float32   `db:"amount"`
	Source     string    `db:"source"`
	Purpose    string    `db:"purpose"`
	AccountID  string    `db:"account_id"`
	Budget     string    `db:"budget"`
	Notes      string    `db:"notes"`
}

// Debit defines a debit in the database
type Debit struct {
	ID         string    `db:"id"`
	PostedDate time.Time `db:"posted_date"`
	Amount     float32   `db:"amount"`
	Vendor     string    `db:"vendor"`
	Purpose    string    `db:"purpose"`
	AccountID  string    `db:"account_id"`
	Budget     string    `db:"budget"`
	Notes      string    `db:"notes"`
}
