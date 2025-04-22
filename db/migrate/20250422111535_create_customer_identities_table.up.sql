-- This is a SQL migration script for creating the customer_identities table
-- This table stores customer identity information
-- such as NIK, Passport, etc.
-- It is linked to the customers table via customer_id
-- and ensures that each customer can have multiple identities
-- but only one of each type (e.g. NIK, Passport, etc.)
CREATE TABLE IF NOT EXISTS customer_identities (
    id BIGSERIAL PRIMARY KEY,                -- Auto-incrementing ID
    customer_id BIGINT NOT NULL,             -- Foreign key to customers table
    identity_type SMALLINT NOT NULL,         -- e.g. 1 = NIK, 2 = Passport, etc.
    identity_number VARCHAR(32) NOT NULL  ,  -- e.g. NIK, Passport Number, etc.
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Automatically set creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Automatically set updated timestamp

    CONSTRAINT uq_customer_identity_type UNIQUE(customer_id, identity_type),      -- prevent duplicates
);
