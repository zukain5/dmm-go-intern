package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Relationship interface {
	// Create relationship
	Create(ctx context.Context, follower *object.Account, followee *object.Account) (int64, bool, error)

	// Whether relationship exists
	Exists(ctx context.Context, follower_id int64, followee_id int64) bool
}
