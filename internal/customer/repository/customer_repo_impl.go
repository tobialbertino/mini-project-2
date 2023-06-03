package repository

import (
	"database/sql"
	entity "miniProject2/internal/customer/model/entoty"
)

type CustomerRepositoryImpl struct{}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

// CreateCustomer implements CustomerRepository.
func (*CustomerRepositoryImpl) CreateCustomer(tx *sql.Tx, et entity.Customer) (int64, error) {
	panic("unimplemented")
}

// DeleteCustomerByID implements CustomerRepository.
func (*CustomerRepositoryImpl) DeleteCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error) {
	panic("unimplemented")
}

// GetAllCustomer implements CustomerRepository.
func (*CustomerRepositoryImpl) GetAllCustomer(tx *sql.Tx, et entity.Customer) ([]entity.Customer, error) {
	panic("unimplemented")
}

// GetCustomerByID implements CustomerRepository.
func (*CustomerRepositoryImpl) GetCustomerByID(tx *sql.Tx, et entity.Customer) (entity.Customer, error) {
	panic("unimplemented")
}

// UpdateCustomerByID implements CustomerRepository.
func (*CustomerRepositoryImpl) UpdateCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error) {
	panic("unimplemented")
}
