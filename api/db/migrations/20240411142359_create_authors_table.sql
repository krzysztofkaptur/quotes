-- +goose Up
create table authors (
  id serial primary key,
  name varchar(50) unique not null 
);

-- +goose Down
drop table authors;
