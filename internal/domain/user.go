package domain

import "github.com/google/uuid"

func NewUser(emailAddress string) User {
	return &UserRaw{
		id:           uuid.New(),
		emailAddress: emailAddress,
	}
}

type User interface {
	ID() uuid.UUID
	EmailAddress() string
}

type UserRaw struct {
	id           uuid.UUID
	emailAddress string
}

func (u *UserRaw) ID() uuid.UUID {
	return u.id
}

func (u *UserRaw) EmailAddress() string {
	return u.emailAddress
}
