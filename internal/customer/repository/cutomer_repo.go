package repository

import (
	"database/sql"
	entity "miniProject2/internal/customer/model/entoty"
)

type CustomerRepository interface {
	GetAllCustomer(tx *sql.Tx, et entity.Customer) ([]entity.Customer, error)
	GetCustomerByID(tx *sql.Tx, et entity.Customer) (entity.Customer, error)

	CreateCustomer(tx *sql.Tx, et entity.Customer) (int64, error)
	UpdateCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error)
	DeleteCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error)
}
