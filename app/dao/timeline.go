package dao

import (
	"context"
	"log"
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
	var statuses *object.Timeline
	err := r.db.SelectContext(ctx, &statuses, `SELECT * FROM status`)
	if err != nil {
		return nil, err
	}
	log.Println(statuses)
	return statuses, nil
}
