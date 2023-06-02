package usecase

import (
	"database/sql"
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

// AddActor implements AccountUseCase.
func (uc *AccountUseCaseImpl) AddActor(req domain.Actor) (string, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return "", err
	}
	defer helper.CommitOrRollback(err, tx)

	// Hash Passwword
	hashPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	data := entity.Actor{

		Username:   req.Username,
		Password:   hashPassword,
		RoleID:     1,
		IsActive:   false,
		IsVerified: false,
	}

	result, err := uc.AccountRepository.AddActor(tx, data)
	if err != nil {
		return "", err
	}

	return result, nil
}
