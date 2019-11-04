CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(45),
    last_name VARCHAR(45),
    email VARCHAR(100) NOT NULL UNIQUE,
    bearer VARCHAR,
    passwordhash VARCHAR(120),
    emailverified BOOLEAN DEFAULT FALSE,
    emailverifiedtoken VARCHAR(45)
);