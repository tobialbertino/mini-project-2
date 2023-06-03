package usecase

import (
	"miniProject2/internal/account/model/domain"
	"miniProject2/internal/account/model/entity"
)

func DTOActor(et entity.Actor) domain.Actor {
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

func DTOActorList(et []entity.Actor) []domain.Actor {
	var result []domain.Actor = make([]domain.Actor, 0)
	for _, v := range et {
		result = append(result, DTOActor(v))
	}

	return result
}

func DTOAdminReg(et entity.AdminReg) domain.AdminReg {
	return domain.AdminReg{
		ID:           et.ID,
		AdminId:      et.AdminId,
		SuperAdminID: et.SuperAdminID,
		Status:       et.Status,
	}
}

func DTOListAdminReg(et []entity.AdminReg) []domain.AdminReg {
	result := make([]domain.AdminReg, 0)
	for _, v := range et {
		result = append(result, DTOAdminReg(v))
	}

	return result
}
