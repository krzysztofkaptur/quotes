// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: quotes.sql

package database

import (
	"context"
)

const fetchQuotes = `-- name: FetchQuotes :many
select id, text, author_id from quotes
`

func (q *Queries) FetchQuotes(ctx context.Context) ([]Quote, error) {
	rows, err := q.db.QueryContext(ctx, fetchQuotes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quote
	for rows.Next() {
		var i Quote
		if err := rows.Scan(&i.ID, &i.Text, &i.AuthorID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchRandomQuote = `-- name: FetchRandomQuote :one
select quotes.id, text, author_id, authors.id, name 
from quotes
join authors on authors.id = quotes.author_id
order by random() 
limit 1
`

type FetchRandomQuoteRow struct {
	ID       int32  `json:"id"`
	Text     string `json:"text"`
	AuthorID int32  `json:"author_id"`
	ID_2     int32  `json:"id_2"`
	Name     string `json:"name"`
}

func (q *Queries) FetchRandomQuote(ctx context.Context) (FetchRandomQuoteRow, error) {
	row := q.db.QueryRowContext(ctx, fetchRandomQuote)
	var i FetchRandomQuoteRow
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.AuthorID,
		&i.ID_2,
		&i.Name,
	)
	return i, err
}
