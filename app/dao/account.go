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
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

// Create : ユーザの作成
func (r *account) Create(ctx context.Context, a *object.Account) (*object.Account, error) {
	res, err := r.db.NamedExecContext(
		ctx,
		`INSERT INTO account (
			username, password_hash
		) VALUES (
			:username, :password_hash
		)`,
		&a,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create account into db: %w", err)
	}

	// TODO: エラー時の対応
	id, _ := res.LastInsertId()
	entity, _ := r.find(ctx, id)
	return entity, nil
}

// id から オブジェクトを見つけるプライベートメソッド
func (r *account) find(ctx context.Context, id int64) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, `SELECT * FROM account WHERE id = ?`, id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}
