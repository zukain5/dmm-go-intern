package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Relationship interface {
	// Create relationship
	Create(ctx context.Context, follower *object.Account, followee *object.Account) (int64, bool, error)
}
