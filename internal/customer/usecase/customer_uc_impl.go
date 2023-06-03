package usecase

import (
	"database/sql"
	"miniProject2/internal/customer/model/domain"
	"miniProject2/internal/customer/model/entity"
	"miniProject2/internal/customer/repository"
	"miniProject2/pkg/helper"
)

type CustomerUseCaseImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
}

func NewCustomerRepository(CustomerRepo repository.CustomerRepository, DB *sql.DB) CustomertUseCase {
	return &CustomerUseCaseImpl{
		CustomerRepository: CustomerRepo,
		DB:                 DB,
	}
}

// CreateCustomer implements CustomertUseCase.
func (uc *CustomerUseCaseImpl) CreateCustomer(dt domain.Customer) (int64, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	et := entity.Customer{
		FirstName: dt.FirstName,
		LastName:  dt.LastName,
		Email:     dt.Email,
		Avatar:    dt.Avatar,
	}
	i, err := uc.CustomerRepository.CreateCustomer(tx, et)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// DeleteCustomerByID implements CustomertUseCase.
func (uc *CustomerUseCaseImpl) DeleteCustomerByID(dt domain.Customer) (int64, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	et := entity.Customer{
		ID: dt.ID,
	}
	i, err := uc.CustomerRepository.DeleteCustomerByID(tx, et)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// GetAllCustomer implements CustomertUseCase.
func (uc *CustomerUseCaseImpl) GetAllCustomer(dt domain.Customer) ([]domain.Customer, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return []domain.Customer{}, err
	}
	defer helper.CommitOrRollback(err, tx)

	et := entity.Customer{
		FirstName: dt.FirstName,
		LastName:  dt.LastName,
		Email:     dt.Email,
	}
	res, err := uc.CustomerRepository.GetAllCustomer(tx, et)
	if err != nil {
		return []domain.Customer{}, err
	}

	return DTOListCustomer(res), nil
}

// GetCustomerByID implements CustomertUseCase.
func (uc *CustomerUseCaseImpl) GetCustomerByID(dt domain.Customer) (domain.Customer, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return domain.Customer{}, err
	}
	defer helper.CommitOrRollback(err, tx)

	et := entity.Customer{
		ID: dt.ID,
	}
	res, err := uc.CustomerRepository.GetCustomerByID(tx, et)
	if err != nil {
		return domain.Customer{}, err
	}

	return DTOCustomer(res), nil
}

// UpdateCustomerByID implements CustomertUseCase.
func (uc *CustomerUseCaseImpl) UpdateCustomerByID(dt domain.Customer) (int64, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	et := entity.Customer{
		ID: dt.ID,
	}
	res, err := uc.CustomerRepository.DeleteCustomerByID(tx, et)
	if err != nil {
		return 0, err
	}

	return res, nil
}
