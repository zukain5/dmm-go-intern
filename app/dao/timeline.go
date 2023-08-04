package dao

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Timeline
	timeline struct {
		db *sqlx.DB
	}
)

// Create timeline repository
func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func (r *timeline) FindPublicTimelines(ctx context.Context, p repository.FindPublicTimelinesParams) (*object.Timeline, error) {
	var statuses []struct {
		*object.Status
		*object.Account `db:"account"`
	}

	query := `
		SELECT
			s.*,
			a.id "account.id",
			a.username "account.username",
			a.display_name "account.display_name",
			a.avatar "account.avatar",
			a.header "account.header",
			a.note "account.note",
			a.create_at "account.create_at"
		FROM
			status AS s
			JOIN account AS a
				ON s.account_id = a.id
		WHERE
			s.id >= ?
			AND s.id <= ?
		LIMIT
			?
	`
	err := r.db.SelectContext(ctx, &statuses, query, 2, 4, 2)
	if err != nil {
		return nil, err
	}

	var t object.Timeline
	for _, s := range statuses {
		s.Status.Account = s.Account
		t = append(t, s.Status)
	}
	return &t, nil
}
