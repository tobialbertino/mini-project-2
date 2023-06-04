package usecase

import "miniProject2/internal/customer/model/domain"

type CustomertUseCase interface {
	GetAllCustomer(dt domain.Customer, pagi domain.Pagiantion) ([]domain.Customer, domain.Pagiantion, error)
	GetCustomerByID(dt domain.Customer) (domain.Customer, error)

	CreateCustomer(dt domain.Customer) (int64, error)
	UpdateCustomerByID(dt domain.Customer) (int64, error)
	DeleteCustomerByID(dt domain.Customer) (int64, error)
}
