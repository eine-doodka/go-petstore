CREATE TABLE users(
    id bigserial primary key,
    email varchar not null unique,
    crypto varchar not null
);