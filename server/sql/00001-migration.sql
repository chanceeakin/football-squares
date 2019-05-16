
-- Running migration
ALTER TABLE users ADD COLUMN password TEXT;
ALTER TABLE users ADD COLUMN is_admin BOOLEAN;

-- Migration 00001 finished
