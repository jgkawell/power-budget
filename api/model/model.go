package model

type Config struct {
	Env            string
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

// ACCOUNT

const CreateAccount = `
	INSERT INTO accounts(
		id,
		name,
		balance,
		total_in,
		total_out,
		type,
		card_number,
		account_number)
	VALUES(
		:id,
		:name,
		:balance,
		:total_in,
		:total_out,
		:type,
		:card_number,
		:account_number)
	RETURNING *;`

const ReadAccount = `
	SELECT *
	FROM accounts
	WHERE id = $1;`

const UpdateAccount = `
	UPDATE accounts
	SET name = :name,
		balance = :balance,
		total_in = :total_in,
		total_out = :total_out,
		type = :type,
		card_number = :card_number,
		account_number = :account_number
	WHERE id = :id;`

const DeleteAccount = `
	DELETE FROM accounts
	WHERE id = $1;`

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
