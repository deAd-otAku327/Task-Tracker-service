CREATE TABLE IF NOT EXISTS dashboards
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);