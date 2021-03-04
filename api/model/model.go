package model

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
