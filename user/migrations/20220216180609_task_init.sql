-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table task
(
    id           uuid default uuid_generate_v4() not null
        constraint task_pkey
            primary key,
    user_id      uuid                            not null,
    time_created timestamp                       not null,
    time_updated timestamp                       not null,
    header       varchar(256)                    not null,
    description  text
);

create index user_index
    on task (user_id);

alter table task
    owner to postgres;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
