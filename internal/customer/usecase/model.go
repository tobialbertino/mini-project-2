package usecase

import (
	"miniProject2/internal/customer/model/domain"
	"miniProject2/internal/customer/model/entity"
)

func DTOCustomer(et entity.Customer) domain.Customer {
	return domain.Customer{
		ID:        et.ID,
		FirstName: et.FirstName,
		LastName:  et.LastName,
		Email:     et.Email,
		Avatar:    et.Avatar,
		CreatedAt: et.CreatedAt,
		UpdatedAt: et.UpdatedAt,
	}
}

func DTOListCustomer(et []entity.Customer) []domain.Customer {
	result := make([]domain.Customer, 0)
	for _, v := range et {
		result = append(result, DTOCustomer(v))
	}

	return result
}
