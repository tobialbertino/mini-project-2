package domain

import "time"

type Customer struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ListActorWithPaging struct {
	Pagination
	Customers []Customer
}
