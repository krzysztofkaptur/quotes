-- name: FetchQuotes :many
select * from quotes;

-- name: FetchRandomQuote :one
select * 
from quotes
join authors on authors.id = quotes.author_id
order by random() 
limit 1;