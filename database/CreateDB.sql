CREATE TABLE debits (
    id UUID PRIMARY KEY,
    posted_date DATE NOT NULL,
    amount NUMERIC(12, 2) NOT NULL,
    vendor TEXT NOT NULL,
    purpose TEXT NOT NULL,
    account_id UUID NOT NULL,
    budget SMALLINT NOT NULL,
    notes TEXT NULL
);

CREATE TABLE credits (
    id UUID PRIMARY KEY,
    posted_date DATE NOT NULL,
    amount NUMERIC(12, 2) NOT NULL,
    source TEXT NOT NULL,
    purpose TEXT NOT NULL,
    account_id UUID NOT NULL,
    budget SMALLINT NOT NULL,
    notes TEXT NULL
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    balance NUMERIC(12, 2) NOT NULL,
    total_in NUMERIC(12, 2) NOT NULL,
    total_out NUMERIC(12, 2) NOT NULL,
    type TEXT NOT NULL,
    card_number TEXT NULL,
    account_number TEXT NULL
);
