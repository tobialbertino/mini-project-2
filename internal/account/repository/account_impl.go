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

// Pagination implements CustomerRepository.
func (repo *AccountRepositoryImpl) Pagination(tx *sql.Tx, et entity.Pagiantion) (entity.Pagiantion, error) {
	var res entity.Pagiantion

	SQL := `
	SELECT count(id) 
	FROM actors
	WHERE role_id = 1
	`
	varArgs := []interface{}{}

	rows, err := tx.Query(SQL, varArgs...)
	if err != nil {
		return entity.Pagiantion{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&res.Total)
		if err != nil {
			return entity.Pagiantion{}, err
		}
	}

	return res, nil
}

// GetAllAdmin implements AccountRepository.
func (repo *AccountRepositoryImpl) GetAllAdmin(tx *sql.Tx, actor entity.Actor, etPage entity.Pagiantion) ([]entity.Actor, error) {
	result := make([]entity.Actor, 0)

	SQL := `
	SELECT id, username, role_id, is_verified, is_active, created_at, updated_at
	FROM actors
	WHERE LOWER(username) LIKE ?
	AND role_id = 1
	LIMIT ?, ?`
	varArgs := []interface{}{
		fmt.Sprintf("%%%s%%", actor.Username),
		etPage.Offset,
		etPage.PerPage,
	}

	rows, err := tx.Query(SQL, varArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res entity.Actor
	for rows.Next() {
		err := rows.Scan(&res.ID, &res.Username, &res.RoleID, &res.IsVerified, &res.IsActive, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return result, nil
}

// DeleteAdminByID implements AccountRepository.
func (repo *AccountRepositoryImpl) DeleteAdminByID(tx *sql.Tx, actor entity.Actor) (int64, error) {
	SQL := `
	DELETE FROM
		actors
	WHERE
		id = ?`
	varArgs := []interface{}{
		actor.ID,
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

// UpdateAdminStatusByAdminID implements AccountRepository.
func (repo *AccountRepositoryImpl) UpdateAdminStatusByAdminID(tx *sql.Tx, actor entity.Actor) (int64, error) {
	SQL := `
	UPDATE actors 
	SET is_verified=?, is_active=? 
	WHERE id = ?`
	varArgs := []interface{}{
		actor.IsVerified,
		actor.IsActive,
		actor.ID,
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

// UpdateAdminRegStatusByAdminID implements AccountRepository.
func (repo *AccountRepositoryImpl) UpdateAdminRegStatusByAdminID(tx *sql.Tx, adminReg entity.AdminReg) (int64, error) {
	SQL := `
	UPDATE admin_reg 
	SET status=?
	WHERE admin_id = ?`
	varArgs := []interface{}{
		adminReg.Status,
		adminReg.AdminId,
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

// GetAllApprovalAdmin implements AccountRepository.
func (repo *AccountRepositoryImpl) GetAllApprovalAdmin(tx *sql.Tx) ([]entity.AdminReg, error) {
	result := make([]entity.AdminReg, 0)

	SQL := `
	SELECT id, admin_id, super_admin_id, status
	FROM admin_reg `

	rows, err := tx.Query(SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res entity.AdminReg
	for rows.Next() {
		err := rows.Scan(&res.ID, &res.AdminId, &res.SuperAdminID, &res.Status)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return result, nil
}

// VerifyActorCredential implements AccountRepository.
func (repo *AccountRepositoryImpl) VerifyActorCredential(tx *sql.Tx, actor entity.Actor) (entity.Actor, error) {
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
