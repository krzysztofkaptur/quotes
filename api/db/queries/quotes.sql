-- name: FetchQuotes :many
select * from quotes;

-- name: FetchRandomQuote :one
select * 
from quotes
join authors on authors.id = quotes.author_id
order by random() 
limit 1;

-- name: CreateQuote :exec
insert into quotes (author_id, text) values ($1, $2);