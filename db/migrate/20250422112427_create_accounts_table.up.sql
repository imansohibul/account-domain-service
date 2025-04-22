
-- This SQL script creates the accounts table in the database.
-- The table is designed to store account information related to customers.
CREATE TABLE IF NOT EXISTS accounts (
    id BIGSERIAL PRIMARY KEY,                       -- Auto-incrementing ID
    customer_id BIGINT NOT NULL,                    -- Foreign key to customers table
    account_number VARCHAR(16) NOT NULL,            -- Unique account number
    account_type SMALLINT NOT NULL,                 -- e.g., 1 = Savings (tabungan)
    status SMALLINT NOT NULL DEFAULT 1,             -- 1 = Active
    balance NUMERIC(15, 2) NOT NULL DEFAULT 0,      -- Account balance
    currency SMALLINT NOT NULL DEFAULT 1,           -- Currency code (ISO 4217) e.g., 1 = IDR
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Automatically set creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Automatically set updated timestamp

    CONSTRAINT uq_account_number UNIQUE(account_number) -- Prevent duplicates
    INDEX idx_accounts_customer_id ON accounts(customer_id) -- Index for quick lookup by customer_id
);
