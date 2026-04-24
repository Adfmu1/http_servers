-- name: GetUserFromRefreshToken :one
SELECT * from users
WHERE id = (
    SELECT user_id 
    FROM refresh_tokens
    WHERE token = $1
);