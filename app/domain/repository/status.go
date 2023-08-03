package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Create status
	Create(ctx context.Context, s *object.Status) error

	// Find status
	Find(ctx context.Context, id string) (*object.Status, error)
}
