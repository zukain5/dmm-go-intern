package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
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
func (r *status) Create(ctx context.Context, s *object.Status) (*object.Status, error) {
	res, err := r.db.NamedExecContext(
		ctx,
		`INSERT INTO status (
			content, account_id
		) VALUES (
			:content, :account_id
		)`,
		map[string]interface{}{
			"content":    s.Content,
			"account_id": s.Account.ID,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create status into db: %w", err)
	}

	// TODO: エラー対応
	id, _ := res.LastInsertId()
	entity, _ := r.Find(ctx, strconv.FormatInt(id, 10))

	return entity, nil
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
