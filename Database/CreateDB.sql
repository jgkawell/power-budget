CREATE TABLE debits (
    id SERIAL PRIMARY KEY,
    posted_date DATE NOT NULL,
    amount NUMERIC(12, 2) NOT NULL,
    vendor TEXT NOT NULL,
    purpose TEXT NOT NULL,
    account TEXT NOT NULL,
    budget SMALLINT NOT NULL,
    notes TEXT NULL
);

CREATE TABLE credits (
    id SERIAL PRIMARY KEY,
    posted_date DATE NOT NULL,
    amount NUMERIC(12, 2) NOT NULL,
    source TEXT NOT NULL,
    purpose TEXT NOT NULL,
    account TEXT NOT NULL,
    budget SMALLINT NOT NULL,
    notes TEXT NULL
);

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    balance NUMERIC(12, 2) NOT NULL,
    total_in NUMERIC(12, 2) NOT NULL,
    total_out NUMERIC(12, 2) NOT NULL,
    type TEXT NOT NULL,
    card_number TEXT NULL,
    account_number TEXT NULL
);
