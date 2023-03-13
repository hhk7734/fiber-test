package application

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hhk7734/fiber-test/internal/domain"
	"github.com/hhk7734/go-optional"
)

type UpdateToDo struct {
	ID          uuid.UUID
	Title       optional.Optional[string]
	Category    optional.Optional[*string]
	Description optional.Optional[string]
	DueDate     optional.Optional[*time.Time]
	FinishedAt  optional.Optional[*time.Time]
}

type ToDoService interface {
	Create(ctx context.Context, userID uuid.UUID, title string, category *string, description string, dueDate *time.Time) (domain.ToDo, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.ToDo, error)
	GetListByUserID(ctx context.Context, userID uuid.UUID, p *domain.Pagination) ([]domain.ToDo, error)
	Update(ctx context.Context, update UpdateToDo) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
