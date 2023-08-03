package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// Create : ステータスの作成
func (r *status) Create(ctx context.Context, s *object.Status) error {
	_, err := r.db.NamedExecContext(
		ctx,
		`INSERT INTO status (
			content, create_at, account_id
		) VALUES (
			:content, :create_at, :account_id
		)`,
		map[string]interface{}{
			"content":    s.Content,
			"create_at":  s.CreateAt,
			"account_id": s.Account.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create status into db: %w", err)
	}
	return nil
}
