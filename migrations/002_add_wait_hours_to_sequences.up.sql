-- 002_add_wait_hours_to_sequences.up.sql
ALTER TABLE sequences
ADD COLUMN wait_hours INT DEFAULT 24;
