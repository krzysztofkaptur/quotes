-- +goose Up
create table quotes (
  id serial primary key,
  text text not null,
  author_id integer references authors(id) not null,
  unique(text, author_id)
);

-- +goose Down
drop table quotes;
