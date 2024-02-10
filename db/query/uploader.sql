-- name: CreateAdmin :exec
INSERT INTO "uploader" (
    username,
    password
) VALUES ($1, $2);

