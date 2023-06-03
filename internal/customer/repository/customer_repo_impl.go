package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"miniProject2/internal/customer/model/entity"
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
	SQL := `
	DELETE FROM
		customers
	WHERE
		id = ?`
	varArgs := []interface{}{
		et.ID,
	}

	result, err := tx.Exec(SQL, varArgs...)
	if err != nil {
		return 0, err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return i, nil
}

// GetAllCustomer implements CustomerRepository.
func (repo *CustomerRepositoryImpl) GetAllCustomer(tx *sql.Tx, et entity.Customer) ([]entity.Customer, error) {
	result := make([]entity.Customer, 0)

	SQL := `
	SELECT id, first_name, last_name, email, avatar, created_at, updated_at
	FROM customers
	WHERE LOWER(first_name) LIKE ?
	OR LOWER(lastName) LIKE ?
	AND LOWER(email) like ?`
	varArgs := []interface{}{
		fmt.Sprintf("%%%s%%", et.FirstName),
		fmt.Sprintf("%%%s%%", et.LastName),
		fmt.Sprintf("%%%s%%", et.Email),
	}

	rows, err := tx.Query(SQL, varArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res entity.Customer
	for rows.Next() {
		err := rows.Scan(&res.ID, &res.FirstName, &res.LastName, &res.Email, &res.Avatar, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return result, nil
}

// GetCustomerByID implements CustomerRepository.
func (repo *CustomerRepositoryImpl) GetCustomerByID(tx *sql.Tx, et entity.Customer) (entity.Customer, error) {
	res := entity.Customer{}

	SQL := `
	SELECT id, first_name, last_name, email, avatar, created_at, updated_at
	FROM customers
	WHERE id = ?`
	varArgs := []interface{}{
		et.ID,
	}

	rows, err := tx.Query(SQL, varArgs...)
	if err != nil {
		return entity.Customer{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&res.ID, &res.FirstName, &res.LastName, &res.Email, &res.Avatar, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return entity.Customer{}, err
		}
	} else {
		return entity.Customer{}, errors.New("customer Not Found")
	}

	return res, nil
}

// UpdateCustomerByID implements CustomerRepository.
func (repo *CustomerRepositoryImpl) UpdateCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error) {
	SQL := `
	DELETE FROM
		customers
	WHERE
		id = ?`
	varArgs := []interface{}{
		et.ID,
	}

	result, err := tx.Exec(SQL, varArgs...)
	if err != nil {
		return 0, err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return i, nil
}
