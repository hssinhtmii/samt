// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: uploader.sql

package db

import (
	"context"
)

const createAdmin = `-- name: CreateAdmin :exec
INSERT INTO "uploader" (
    username,
    password
) VALUES ($1, $2)
`

type CreateAdminParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) CreateAdmin(ctx context.Context, arg CreateAdminParams) error {
	_, err := q.db.ExecContext(ctx, createAdmin, arg.Username, arg.Password)
	return err
}