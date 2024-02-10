-- name: CreateSessions :one
INSERT INTO "sessions" (
        id,
        phone_number,
        refresh_token,
        is_blocked,
        expired_at
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetSessions :one
SELECT *
FROM "sessions"
WHERE id = $1;