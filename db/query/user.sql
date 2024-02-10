-- name: CreateUser :exec
INSERT INTO "user" (
        first_name,
        last_name,
        phone_number,
        email,
        password
    )
VALUES ($1, $2, $3, $4, $5);

-- name: GetAllUsers :many
SELECT first_name, last_name, phone_number, email
FROM "user";

-- name: GetUser :one
SELECT *
from "user"
WHERE phone_number = $1;

-- name: LoginUser :one
SELECT first_name, last_name, phone_number, email
from "user"
WHERE phone_number = $1 AND password = $2;

-- name: DeleteUser :exec
DELETE FROM "user" WHERE phone_number = $1;

-- name: UpdateUser :one
UPDATE "user" 
SET 
first_name = COALESCE(sqlc.narg(first_name), first_name),
last_name = COALESCE(sqlc.narg(last_name), last_name),
email = COALESCE(sqlc.narg(email), email),
password = COALESCE(sqlc.narg(password), password)
WHERE 
phone_number = $1
RETURNING first_name, last_name, phone_number, email;