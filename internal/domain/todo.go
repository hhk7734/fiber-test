package domain

import (
	"time"

	"github.com/google/uuid"
)

func NewToDo(
	userID uuid.UUID, title string, category *string, description string,
	dueDate *time.Time,
) ToDo {
	return &ToDoRaw{
		id:          uuid.New(),
		userID:      userID,
		title:       title,
		category:    category,
		description: description,
		dueDate:     dueDate,
		finishedAt:  nil,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
		deletedAt:   nil,
	}
}

type ToDo interface {
	ID() uuid.UUID
	UserID() uuid.UUID
	Title() string
	SetTitle(string) error
	Category() *string
	SetCategory(*string) error
	Description() string
	SetDescription(string) error
	DueDate() *time.Time
	SetDueDate(*time.Time) error
	FinishedAt() *time.Time
	SetFinishedAt(*time.Time) error
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time
}

type ToDoRaw struct {
	id          uuid.UUID
	userID      uuid.UUID
	title       string
	category    *string
	description string
	dueDate     *time.Time
	finishedAt  *time.Time
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

func (t *ToDoRaw) ID() uuid.UUID {
	return t.id
}

func (t *ToDoRaw) UserID() uuid.UUID {
	return t.userID
}

func (t *ToDoRaw) Title() string {
	return t.title
}

func (t *ToDoRaw) SetTitle(title string) error {
	t.title = title
	return nil
}

func (t *ToDoRaw) Category() *string {
	return t.category
}

func (t *ToDoRaw) SetCategory(category *string) error {
	t.category = category
	return nil
}

func (t *ToDoRaw) Description() string {
	return t.description
}

func (t *ToDoRaw) SetDescription(description string) error {
	t.description = description
	return nil
}

func (t *ToDoRaw) DueDate() *time.Time {
	return t.dueDate
}

func (t *ToDoRaw) SetDueDate(dueDate *time.Time) error {
	t.dueDate = dueDate
	return nil
}

func (t *ToDoRaw) FinishedAt() *time.Time {
	return t.finishedAt
}

func (t *ToDoRaw) SetFinishedAt(finishedAt *time.Time) error {
	t.finishedAt = finishedAt
	return nil
}

func (t *ToDoRaw) CreatedAt() time.Time {
	return t.createdAt
}

func (t *ToDoRaw) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *ToDoRaw) DeletedAt() *time.Time {
	return t.deletedAt
}
