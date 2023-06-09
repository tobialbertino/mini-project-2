package repository

import (
	"database/sql"
	"miniProject2/modules/account/model/entity"
)

type AccountRepository interface {
	// auth
	VerifyActorCredential(tx *sql.Tx, actor entity.Actor) (entity.Actor, error)
	Login(tx *sql.Tx, token entity.Token) (string, error) // TODO: store token

	// actor
	AddActor(tx *sql.Tx, actor entity.Actor) (int64, error)
	GetAllAdmin(tx *sql.DB, actor entity.Actor, et entity.Pagination) ([]entity.Actor, error)
	Pagination(tx *sql.DB, et entity.Pagination) (entity.Pagination, error) // only Get Total Data

	// admin reg
	RegisterAdmin(tx *sql.Tx, adminReg entity.AdminReg) (int64, error)

	// super_admin only
	GetAllApprovalAdmin(tx *sql.Tx) ([]entity.AdminReg, error)
	UpdateAdminRegStatusByAdminID(tx *sql.Tx, adminReg entity.AdminReg) (int64, error)
	UpdateAdminStatusByAdminID(tx *sql.Tx, actor entity.Actor) (int64, error)
	DeleteAdminByID(tx *sql.Tx, actor entity.Actor) (int64, error)
}
