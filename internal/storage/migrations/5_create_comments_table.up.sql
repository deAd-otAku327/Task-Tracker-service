CREATE TABLE IF NOT EXISTS comments
(
    id SERIAL PRIMARY KEY,
    task_id INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    date_time TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE comments ADD CONSTRAINT comments_task_id_fk
FOREIGN KEY (task_id) REFERENCES tasks(id);

ALTER TABLE comments ADD CONSTRAINT comments_author_id_fk
FOREIGN KEY (author_id) REFERENCES users(id);