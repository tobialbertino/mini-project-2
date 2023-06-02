package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"miniProject2/internal/account/model/entity"
)

type AccountRepositoryImpl struct{}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{}
}

// VerifyActorCredential implements AccountRepository.
func (*AccountRepositoryImpl) VerifyActorCredential(tx *sql.Tx, actor entity.Actor) (entity.Actor, error) {
	SQL := `
	SELECT id, username, password, role_id, is_verified, is_active, created_at, updated_at
	FROM actors 
	WHERE username = ?`
	varArgs := []interface{}{
		actor.Username,
	}

	rows, err := tx.Query(SQL, varArgs...)
	if err != nil {
		return entity.Actor{}, err
	}
	defer rows.Close()

	res := entity.Actor{}
	if rows.Next() {
		err := rows.Scan(&res.ID, &res.Username, &res.Password, &res.RoleID, &res.IsVerified, &res.IsActive, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return entity.Actor{}, err
		}
	} else {
		return entity.Actor{}, errors.New("incorrect username or password")
	}

	return res, nil
}

// RegisterAdmin implements AccountRepository.
func (repo *AccountRepositoryImpl) RegisterAdmin(tx *sql.Tx, adminReg entity.AdminReg) (int64, error) {
	SQL := `
	INSERT INTO admin_reg(admin_id, super_admin_id, status) 
	VALUES (?, 1, ?)`
	varArgs := []interface{}{
		adminReg.AdminId,
		adminReg.Status,
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

// AddActor implements AccountRepository.
func (repo *AccountRepositoryImpl) AddActor(tx *sql.Tx, actor entity.Actor) (int64, error) {
	SQL := `
	INSERT INTO actors(username, password, role_id, is_active, is_verified) 
	VALUES (?, ?, ?, ?, ?)`
	varArgs := []interface{}{
		actor.Username,
		actor.Password,
		actor.RoleID,
		actor.IsActive,
		actor.IsVerified,
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

// Login implements AccountRepository.
func (repo *AccountRepositoryImpl) Login(tx *sql.Tx, token entity.Token) (string, error) {
	SQL := `INSERT INTO authentications(token) VALUES (?)`
	varArgs := []interface{}{
		token.Token,
	}

	result, err := tx.Exec(SQL, varArgs...)
	if err != nil {
		return "error repository", err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return "error repository", err
	}

	return fmt.Sprintf("rows affected: %d", i), nil
}
