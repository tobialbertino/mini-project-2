package usecase

import (
	"miniProject2/internal/account/model/domain"
	"miniProject2/internal/account/model/entity"
)

func DTO(et entity.Actor) domain.Actor {
	return domain.Actor{
		ID:         et.ID,
		Username:   et.Username,
		Password:   et.Password,
		RoleID:     et.RoleID,
		IsVerified: et.IsVerified,
		IsActive:   et.IsActive,
		CreatedAt:  et.CreatedAt,
		UpdatedAt:  et.UpdatedAt,
	}
}
