-- +goose Up
-- +goose StatementBegin
create table users (
    id serial primary key,
    login varchar(255) not null unique,
    password varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
