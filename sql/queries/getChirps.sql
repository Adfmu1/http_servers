-- name: GetChirps :many
SELECT * FROM chirps
ORDER BY CREATED_AT ASC;