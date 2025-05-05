CREATE TYPE task_status_enum AS ENUM (
    'created',
    'in_progress',
    'done',
    'dropped'
);

CREATE TABLE IF NOT EXISTS tasks
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    discription TEXT,
    status task_status_enum NOT NULL DEFAULT 'in_progress',
    assignie_id INTEGER REFERENCES users(id),
    board_id INTEGER REFERENCES dashboards(id) ON DELETE SET NULL,
    updated_ad TIMESTAMP NOT NULL DEFAULT NOW()
);