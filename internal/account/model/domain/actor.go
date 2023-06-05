package domain

import "time"

type Actor struct {
	ID         int64
	Username   string
	Password   string
	RoleID     int64
	IsVerified bool
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ListActorWithPaging struct {
	Pagination
	Admins []Actor
}
