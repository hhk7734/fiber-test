package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Add(ctx context.Context, u User) error
	TakeByID(ctx context.Context, id uuid.UUID) (User, error)
	TakeByEmailAddress(ctx context.Context, emailAddress string) (User, error)
	Update(ctx context.Context, u User) error
	Delete(ctx context.Context, u User) error
}
