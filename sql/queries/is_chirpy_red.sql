-- name: IsChirpyRed :one
SELECT is_chirpy_red FROM users WHERE id = $1;

-- name: SetIsChirpyRed :one
UPDATE users SET is_chirpy_red = $2 WHERE id = $1
RETURNING *;