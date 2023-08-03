package dao

import (
	"context"
	"database/sql"
	"errors"
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

// Find : DB からステータスを見つける
func (r *status) Find(ctx context.Context, id string) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, `SELECT * FROM status WHERE id = ?`, id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	account := new(object.Account)
	err = r.db.
		QueryRowxContext(ctx, `SELECT * FROM account WHERE id = ?`, entity.AccountId).
		StructScan(account)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("associated account for the status not found: %w", err)
		}

		return nil, fmt.Errorf("failed to find associated account for the status: %w", err)
	}

	entity.Account = account
	return entity, nil
}
