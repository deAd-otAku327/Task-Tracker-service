CREATE TABLE IF NOT EXISTS dashboards
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    updated_ad TIMESTAMP NOT NULL DEFAULT NOW()
);