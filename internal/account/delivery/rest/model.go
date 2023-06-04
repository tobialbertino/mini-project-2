package rest

import (
	"miniProject2/internal/account/model/domain"
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

type ReqAddActor struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ReqIDActor struct {
	ID int64 `json:"id" binding:"required"`
}

type ReqUpdateAdminStatus struct {
	AdminID    int64  `json:"admin_id" binding:"required"`
	Status     string `json:"status" binding:"required"`
	IsVerified bool   `json:"is_verified" binding:"boolean"`
	IsActive   bool   `json:"is_active" binding:"boolean"`
}

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ResGetAllAdminWithPaging struct {
	Pagination
	Admins []ResponseActor `json:"admin"`
}

type ResponseActor struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	RoleID     int64     `json:"role_id"`
	IsVerified bool      `json:"is_verified"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ResponseAdminReg struct {
	ID           int64  `json:"id"`
	AdminId      int64  `json:"admin_id"`
	SuperAdminID int64  `json:"super_admin_id"`
	Status       string `json:"status"`
}

func ToResponseActor(dt domain.Actor) ResponseActor {
	return ResponseActor{
		ID:         dt.ID,
		Username:   dt.Username,
		RoleID:     dt.RoleID,
		IsVerified: dt.IsVerified,
		IsActive:   dt.IsActive,
		CreatedAt:  dt.CreatedAt,
		UpdatedAt:  dt.UpdatedAt,
	}
}

func ToResponseAdminReg(dt domain.AdminReg) ResponseAdminReg {
	return ResponseAdminReg{
		ID:           dt.ID,
		AdminId:      dt.AdminId,
		SuperAdminID: dt.SuperAdminID,
		Status:       dt.Status,
	}
}

func ResponseListActor(dt []domain.Actor) []ResponseActor {
	result := make([]ResponseActor, 0)
	for _, v := range dt {
		result = append(result, ToResponseActor(v))
	}

	return result
}

func ResponseListAdminReg(dt []domain.AdminReg) []ResponseAdminReg {
	result := make([]ResponseAdminReg, 0)
	for _, v := range dt {
		result = append(result, ToResponseAdminReg(v))
	}

	return result
}
