package usecase

import (
	"database/sql"
	"errors"
	"miniProject2/internal/account/model/domain"
	"miniProject2/internal/account/model/entity"
	"miniProject2/internal/account/repository"
	"miniProject2/pkg/helper"
	"miniProject2/pkg/security"
)

type AccountUseCaseImpl struct {
	AccountRepository repository.AccountRepository
	DB                *sql.DB
}

func NewAccountUseCase(AccountRepository repository.AccountRepository, DB *sql.DB) AccountUseCase {
	return &AccountUseCaseImpl{
		AccountRepository: AccountRepository,
		DB:                DB,
	}
}

// VerifyActorCredential implements AccountUseCase.
func (uc *AccountUseCaseImpl) VerifyActorCredential(req domain.Actor) (domain.Actor, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return domain.Actor{}, err
	}
	defer helper.CommitOrRollback(err, tx)

	entity := entity.Actor{
		Username: req.Username,
		Password: req.Password,
	}
	result, err := uc.AccountRepository.VerifyActorCredential(tx, entity)
	if err != nil {
		return domain.Actor{}, err
	}

	// compare password
	isValid := security.CheckPasswordHash(req.Password, result.Password)
	if !isValid {
		return domain.Actor{}, errors.New("invalid username or password")
	}
	res := DTO(result)
	return res, nil
}

// AddActor implements AccountUseCase.
func (uc *AccountUseCaseImpl) AddActor(req domain.Actor) (int64, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	// Hash Passwword
	hashPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return 0, err
	}

	data := entity.Actor{
		Username:   req.Username,
		Password:   hashPassword,
		RoleID:     1,
		IsActive:   false,
		IsVerified: false,
	}

	resultID, err := uc.AccountRepository.AddActor(tx, data)
	if err != nil {
		return 0, err
	}

	adminReg := entity.AdminReg{
		AdminId:      resultID,
		SuperAdminID: 1,
		Status:       "pending",
	}

	result, err := uc.AccountRepository.RegisterAdmin(tx, adminReg)
	if err != nil {
		return 0, err
	}

	return result, nil
}
