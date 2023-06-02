package repository

import (
	"database/sql"
	"miniProject2/internal/account/model/entity"
)

type AccountRepository interface {
	// auth
	Login(tx *sql.Tx, token entity.Token) (string, error)

	// actor
	AddActor(tx *sql.Tx, actor entity.Actor) (int64, error)
	RegisterAdmin(tx *sql.Tx, adminReg entity.AdminReg) (int64, error)

	// admin reg

}
