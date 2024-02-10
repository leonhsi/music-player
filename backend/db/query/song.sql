-- name: CreateSong :one
INSERT INTO songs (
  song_name, artist_id, thumbnail_s3_path, mp3_s3_path
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetSong :one
SELECT * FROM songs
WHERE song_id = $1 LIMIT 1;

-- name: ListSongs :many
SELECT * FROM songs
ORDER BY song_id LIMIT $1;

-- name: UpdateSong :one
UPDATE songs
  set song_name = $2,
  artist_id = $3,
  thumbnail_s3_path = $4,
  mp3_s3_path = $5
WHERE song_id = $1
RETURNING *;

-- name: DeleteSong :exec
DELETE FROM songs
WHERE song_id = $1;
