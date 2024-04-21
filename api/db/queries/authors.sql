-- name: FetchAuthors :many
select * from authors;

-- name: CreateAuthor :exec
insert into authors (name) values ($1);