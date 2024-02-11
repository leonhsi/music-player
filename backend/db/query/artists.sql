-- name: GetArtist :one
SELECT * FROM artists
WHERE artist_name = $1 LIMIT 1;

-- name: ListArtists :many
SELECT * FROM artists 
ORDER BY artist_id LIMIT $1;

-- name: CreateArtist :one
INSERT INTO artists (
 artist_name 
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateArtist :one
UPDATE artists
  set artist_name = $2
WHERE artist_id = $1
RETURNING *;

-- name: DeleteArtist :exec
DELETE FROM artists
WHERE artist_id = $1;
