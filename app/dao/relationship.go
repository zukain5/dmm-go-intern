package dao

import (
	"context"
	"database/sql"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Relationship
	relationship struct {
		db *sqlx.DB
	}
)

// Create relationship repository
func NewRelationship(db *sqlx.DB) repository.Relationship {
	return &relationship{db: db}
}

func (r *relationship) Create(ctx context.Context, follwer *object.Account, followee *object.Account) (int64, bool, error) {
	res, err := r.db.ExecContext(
		ctx,
		`INSERT INTO relationship (
			follower_id, followee_id
		) VALUES (
			?, ?
		)`,
		follwer.ID,
		followee.ID,
	)
	if err != nil {
		return -1, false, fmt.Errorf("failed to follow: %w", err)
	}

	// TODO: エラー時の対応
	id, _ := res.LastInsertId()

	followed_by := r.exists(ctx, followee.ID, follwer.ID)
	return id, followed_by, nil
}

func (r *relationship) exists(ctx context.Context, follower_id, followee_id int64) bool {
	query := `
		SELECT
			id
		FROM
			relationship
		WHERE
			follower_id = ?
			AND followee_id = ?
	`
	var id int
	err := r.db.QueryRowxContext(ctx, query, follower_id, followee_id).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}

		// TODO: エラー対応
		return false
	}

	return true
}
