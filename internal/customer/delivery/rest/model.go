package rest

import (
	"miniProject2/internal/customer/model/domain"
	"time"
)

type WebResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type RowsAffected struct {
	Message      string `json:"message"`
	RowsAffected any    `json:"rows_affected"`
}

type ReqGetAllCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Customer struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ResGetAllCustomerWithPaging struct {
	Pagination
	Customer []Customer `json:"customers"`
}

type ReqAddCustomer struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Avatar    string `json:"avatar" binding:"required"`
}

func ToResponseCustomer(dt domain.Customer) Customer {
	return Customer{
		ID:        dt.ID,
		FirstName: dt.FirstName,
		LastName:  dt.LastName,
		Email:     dt.Email,
		Avatar:    dt.Avatar,
		CreatedAt: dt.CreatedAt,
		UpdatedAt: dt.UpdatedAt,
	}
}

func ToResponseListCustomer(dt []domain.Customer) []Customer {
	result := make([]Customer, 0)
	for _, v := range dt {
		result = append(result, ToResponseCustomer(v))
	}

	return result
}
