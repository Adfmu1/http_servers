-- name: GetChirpsByUserID :many
SELECT * FROM chirps
WHERE user_id = $1 
ORDER BY CREATED_AT ASC;