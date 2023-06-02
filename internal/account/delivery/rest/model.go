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

type ResponseActor struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	RoleID     int64     `json:"role_id"`
	IsVerified bool      `json:"is_verified"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToResponse(dt domain.Actor) ResponseActor {
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
