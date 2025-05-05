CREATE TABLE IF NOT EXISTS dashboards
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    discription TEXT,
    updated_ad TIMESTAMP NOT NULL DEFAULT NOW()
);