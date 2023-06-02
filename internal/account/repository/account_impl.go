package repository

import (
	"database/sql"
	"fmt"
	"miniProject2/internal/account/model/entity"
)

type AccountRepositoryImpl struct{}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{}
}

// AddActor implements AccountRepository.
func (*AccountRepositoryImpl) AddActor(tx *sql.Tx, actor entity.Actor) (string, error) {
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
		return "error repository", err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return "error repository", err
	}

	return fmt.Sprintf("rows affected: %d", i), nil
}

// Login implements AccountRepository.
func (*AccountRepositoryImpl) Login(tx *sql.Tx, token entity.Token) (string, error) {
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
