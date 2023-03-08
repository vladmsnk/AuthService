CREATE TABLE IF NOT EXISTS users(
    id uuid,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(2048),
    number VARCHAR(255),
    CONSTRAINT users_pk  PRIMARY KEY (id, username, email) -- делаю pk на 3 поля, чтобы был индекс по username и email
);