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
    description TEXT,
    status task_status_enum NOT NULL DEFAULT 'created',
    author_id INTEGER,
    assignie_id INTEGER,
    board_id INTEGER,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE tasks ADD CONSTRAINT tasks_author_id_fk
FOREIGN KEY (author_id) REFERENCES users(id);

ALTER TABLE tasks ADD CONSTRAINT tasks_assignie_id_fk
FOREIGN KEY (assignie_id) REFERENCES users(id);

ALTER TABLE tasks ADD CONSTRAINT tasks_board_id_fk
FOREIGN KEY (board_id) REFERENCES dashboards(id) ON DELETE SET NULL;