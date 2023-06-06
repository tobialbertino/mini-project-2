package repository

import (
	"database/sql"
	"miniProject2/internal/customer/model/entity"
)

type CustomerRepository interface {
	// pagination
	Pagination(tx *sql.DB, et entity.Pagination) (entity.Pagination, error) // only Get Total Data
	GetAllCustomer(tx *sql.DB, et entity.Customer, etPaging entity.Pagination) ([]entity.Customer, error)
	// get count total data

	GetCustomerByID(tx *sql.Tx, et entity.Customer) (entity.Customer, error)

	CreateCustomer(tx *sql.Tx, et entity.Customer) (int64, error)
	UpdateCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error)
	DeleteCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error)
}
