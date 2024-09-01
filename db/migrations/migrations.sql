-- 001_create_sequences.up.sql
CREATE TABLE sequences (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    open_tracking_enabled BOOLEAN NOT NULL,
    click_tracking_enabled BOOLEAN NOT NULL
);

CREATE TABLE sequence_steps (
    id SERIAL PRIMARY KEY,
    sequence_id INT REFERENCES sequences(id),
    subject VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    order INT NOT NULL
);
