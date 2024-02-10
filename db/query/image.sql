-- name: UploadImage :exec
INSERT INTO "image" (img, is_poster)
VALUES ($1, $2);
-- name: DeleteImage :exec
DELETE FROM "image"
WHERE img = $1;
-- name: GetAllImage :many
SELECT *
FROM "image"
WHERE "is_poster" = $1;