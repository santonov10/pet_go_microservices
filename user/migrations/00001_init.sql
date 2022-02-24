-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table "user"
(
    id       uuid default uuid_generate_v4() not null
        constraint event_pkey
            primary key,
    login    varchar(256)                    not null,
    password varchar(32)                     not null
);

alter table "user"
    owner to postgres;

create unique index login_key
    on "user" (login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
