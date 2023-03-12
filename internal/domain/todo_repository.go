package domain

import (
	"context"

	"github.com/google/uuid"
)

type ToDoRepository interface {
	Add(ctx context.Context, t ToDo) error
	TakeByID(ctx context.Context, id uuid.UUID) (ToDo, error)
	FindByUserID(ctx context.Context, userID uuid.UUID, p *Pagination) ([]ToDo, error)
	Update(ctx context.Context, t ToDo) error
	Delete(ctx context.Context, t ToDo) error
}
