package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type FindPublicTimelinesParams struct {
	max_id   int
	since_id int
	limit    int
}

type Timeline interface {
	FindPublicTimelines(ctx context.Context, p FindPublicTimelinesParams) (*object.Timeline, error)
}
