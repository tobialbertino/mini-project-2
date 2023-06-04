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

func NewCustomerUseCase(CustomerRepo repository.CustomerRepository, DB *sql.DB) CustomertUseCase {
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
func (uc *CustomerUseCaseImpl) GetAllCustomer(dt domain.Customer, pagi domain.Pagiantion) ([]domain.Customer, domain.Pagiantion, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return []domain.Customer{}, domain.Pagiantion{}, err
	}
	defer helper.CommitOrRollback(err, tx)

	// define pagination
	etPaging := entity.Pagiantion{
		Page:       pagi.Page,
		PerPage:    6,                   // always fix 6 data == LIMIT
		Total:      0,                   // after query
		TotalPages: 0,                   // after query, total / PerPage
		Offset:     (pagi.Page - 1) * 6, // (Page-1) * PerPage
	}

	et := entity.Customer{
		FirstName: dt.FirstName,
		LastName:  dt.LastName,
		Email:     dt.Email,
	}
	// get all customer with pagination
	res, err := uc.CustomerRepository.GetAllCustomer(tx, et, etPaging)
	if err != nil {
		return []domain.Customer{}, domain.Pagiantion{}, err
	}

	// Get Total Data
	resPaging, err := uc.CustomerRepository.Pagination(tx, etPaging)
	if err != nil {
		return []domain.Customer{}, domain.Pagiantion{}, err
	}

	totalPages := resPaging.Total / 6
	if resPaging.Total%6 != 0 {
		totalPages++
	}
	etPaging.Total = resPaging.Total
	etPaging.TotalPages = totalPages

	return DTOListCustomer(res), domain.Pagiantion(etPaging), nil
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
		ID:        dt.ID,
		FirstName: dt.FirstName,
		LastName:  dt.LastName,
		Email:     dt.Email,
		Avatar:    dt.Avatar,
	}
	res, err := uc.CustomerRepository.UpdateCustomerByID(tx, et)
	if err != nil {
		return 0, err
	}

	return res, nil
}
