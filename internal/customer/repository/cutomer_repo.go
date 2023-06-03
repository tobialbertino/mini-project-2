package repository

import (
	"database/sql"
	"miniProject2/internal/customer/model/entity"
)

type CustomerRepository interface {
	GetAllCustomer(tx *sql.Tx, et entity.Customer) ([]entity.Customer, error)
	GetCustomerByID(tx *sql.Tx, et entity.Customer) (entity.Customer, error)

	CreateCustomer(tx *sql.Tx, et entity.Customer) (int64, error)
	UpdateCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error)
	DeleteCustomerByID(tx *sql.Tx, et entity.Customer) (int64, error)
}
