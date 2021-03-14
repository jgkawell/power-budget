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
	ID            string  `db:"id" json:"id"`
	Name          string  `db:"name" json:"name"`
	Balance       float32 `db:"balance" json:"balance"`
	TotalIn       float32 `db:"total_in" json:"totalIn"`
	TotalOut      float32 `db:"total_out" json:"totalOut"`
	Type          string  `db:"type" json:"type"`
	CardNumber    string  `db:"card_number" json:"cardNumber"`
	AccountNumber string  `db:"account_number" json:"accountNumber"`
}

// Credit defines a credit in the database
type Credit struct {
	ID         string    `db:"id" json:"id"`
	PostedDate time.Time `db:"posted_date" json:"postedDate" time_format:"2006-01-02"`
	Amount     float32   `db:"amount" json:"amount"`
	Source     string    `db:"source" json:"source"`
	Purpose    string    `db:"purpose" json:"purpose"`
	AccountID  string    `db:"account_id" json:"accountId"`
	Budget     int8      `db:"budget" json:"budget"`
	Notes      string    `db:"notes" json:"notes"`
}

// Debit defines a debit in the database
type Debit struct {
	ID         string    `db:"id" json:"id"`
	PostedDate time.Time `db:"posted_date" json:"postedDate" time_format:"2006-01-02"`
	Amount     float32   `db:"amount" json:"amount"`
	Vendor     string    `db:"vendor" json:"vendor"`
	Purpose    string    `db:"purpose" json:"purpose"`
	AccountID  string    `db:"account_id" json:"accountId"`
	Budget     int8      `db:"budget" json:"budget"`
	Notes      string    `db:"notes" json:"notes"`
}
