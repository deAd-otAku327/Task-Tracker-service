CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(10) NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_username_unique UNIQUE (username);
ALTER TABLE users ADD CONSTRAINT users_email_unique UNIQUE (email);