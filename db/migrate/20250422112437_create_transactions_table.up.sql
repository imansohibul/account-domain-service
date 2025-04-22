-- This SQL script creates a table named 'transactions' in the database.
-- The table is designed to store transaction records related to accounts.
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,                    -- Auto-incrementing ID
    account_id INT NOT NULL,                  -- Account ID (Foreign Key to reference the account)
    type SMALLINT NOT NULL,                   -- Transaction type (e.g., debit, credit)
    amount DECIMAL(15, 2) NOT NULL,           -- Amount involved in the transaction
    initial_balance DECIMAL(15, 2) NOT NULL,  -- Balance before the transaction
    final_balance DECIMAL(15, 2) NOT NULL,    -- Balance after the transaction
    currency SMALLINT NOT NULL DEFAULT 1,     -- Currency code (ISO 4217) e.g., 1 = IDR
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Automatically set updated timestamp
);
