-- This SQL script creates a table named 'customers' in the database.
-- The table is designed to store customer information.
-- It includes fields for the customer's full name, phone number, and timestamps for creation and updates.
-- The phone number is standardized to E.164 format for international compatibility.
-- The table also includes a unique constraint on the phone number to prevent duplicates.
CREATE TABLE IF NOT EXISTS customers (
    id BIGSERIAL PRIMARY KEY,              -- Auto-incrementing ID
    fullname VARCHAR(255) NOT NULL,        -- Full name of the customer
    phone_number VARCHAR(16),              -- Phone number of the customer, Standardization (E.164 format - International Compatibility)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Automatically set creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Automatically set updated timestamp

    CONSTRAINT uq_customer_phone_number UNIQUE(phone_number)  -- Prevent duplicates
);
