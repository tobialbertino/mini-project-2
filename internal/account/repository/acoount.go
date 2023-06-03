package repository

import (
	"database/sql"
	"miniProject2/internal/account/model/entity"
)

type AccountRepository interface {
	// auth
	VerifyActorCredential(tx *sql.Tx, actor entity.Actor) (entity.Actor, error)
	Login(tx *sql.Tx, token entity.Token) (string, error) // TODO: store token

	// actor
	AddActor(tx *sql.Tx, actor entity.Actor) (int64, error)

	// admin reg
	RegisterAdmin(tx *sql.Tx, adminReg entity.AdminReg) (int64, error)

	// super_admin only
	GetAllApprovalAdmin(tx *sql.Tx) ([]entity.AdminReg, error)
	UpdateAdminRegStatusByAdminID(tx *sql.Tx, adminReg entity.AdminReg) (int64, error)
	UpdateAdminStatusByAdminID(tx *sql.Tx, actor entity.Actor) (int64, error)
	DeleteAdminByID(tx *sql.Tx, actor entity.Actor) (int64, error)
}
