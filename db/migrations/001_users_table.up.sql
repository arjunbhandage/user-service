BEGIN;
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    email      VARCHAR(255) UNIQUE NOT NULL,
    password   VARCHAR(255)        NOT NULL
);
COMMIT;