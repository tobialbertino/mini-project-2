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
func (repo *CustomerRepositoryImpl) CreateCustomer(tx *sql.Tx, et entity.Customer) (int64, error) {
	SQL := `
	INSERT INTO customers(first_name, last_name, email, avatar) 
	VALUES (?, ?, ?, ?)`
	varArgs := []interface{}{
		et.FirstName,
		et.LastName,
		et.Email,
		et.Avatar,
	}

	result, err := tx.Exec(SQL, varArgs...)
	if err != nil {
		return 0, err
	}

	i, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return i, nil
}

// DeleteCustomerByID implements CustomerRepository.
func (repo *CustomerRepositoryImpl) DeleteCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error) {
	panic("unimplemented")
}

// GetAllCustomer implements CustomerRepository.
func (repo *CustomerRepositoryImpl) GetAllCustomer(tx *sql.Tx, et entity.Customer) ([]entity.Customer, error) {
	panic("unimplemented")
}

// GetCustomerByID implements CustomerRepository.
func (repo *CustomerRepositoryImpl) GetCustomerByID(tx *sql.Tx, et entity.Customer) (entity.Customer, error) {
	panic("unimplemented")
}

// UpdateCustomerByID implements CustomerRepository.
func (repo *CustomerRepositoryImpl) UpdateCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error) {
	panic("unimplemented")
}
