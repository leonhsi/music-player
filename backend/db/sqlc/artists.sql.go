// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: artists.sql

package db

import (
	"context"
)

const createArtist = `-- name: CreateArtist :one
INSERT INTO artists (
 artist_name 
) VALUES (
  $1
)
RETURNING artist_id, artist_name
`

func (q *Queries) CreateArtist(ctx context.Context, artistName string) (Artist, error) {
	row := q.db.QueryRowContext(ctx, createArtist, artistName)
	var i Artist
	err := row.Scan(&i.ArtistID, &i.ArtistName)
	return i, err
}

const deleteArtist = `-- name: DeleteArtist :exec
DELETE FROM artists
WHERE artist_id = $1
`

func (q *Queries) DeleteArtist(ctx context.Context, artistID int64) error {
	_, err := q.db.ExecContext(ctx, deleteArtist, artistID)
	return err
}

const getArtist = `-- name: GetArtist :one
SELECT artist_id, artist_name FROM artists
WHERE artist_name = $1 LIMIT 1
`

func (q *Queries) GetArtist(ctx context.Context, artistName string) (Artist, error) {
	row := q.db.QueryRowContext(ctx, getArtist, artistName)
	var i Artist
	err := row.Scan(&i.ArtistID, &i.ArtistName)
	return i, err
}

const listArtists = `-- name: ListArtists :many
SELECT artist_id, artist_name FROM artists 
ORDER BY artist_id LIMIT $1
`

func (q *Queries) ListArtists(ctx context.Context, limit int32) ([]Artist, error) {
	rows, err := q.db.QueryContext(ctx, listArtists, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Artist{}
	for rows.Next() {
		var i Artist
		if err := rows.Scan(&i.ArtistID, &i.ArtistName); err != nil {
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

const updateArtist = `-- name: UpdateArtist :one
UPDATE artists
  set artist_name = $2
WHERE artist_id = $1
RETURNING artist_id, artist_name
`

type UpdateArtistParams struct {
	ArtistID   int64  `json:"artist_id"`
	ArtistName string `json:"artist_name"`
}

func (q *Queries) UpdateArtist(ctx context.Context, arg UpdateArtistParams) (Artist, error) {
	row := q.db.QueryRowContext(ctx, updateArtist, arg.ArtistID, arg.ArtistName)
	var i Artist
	err := row.Scan(&i.ArtistID, &i.ArtistName)
	return i, err
}