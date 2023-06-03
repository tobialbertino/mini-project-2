package usecase

import (
	"database/sql"
	"errors"
	"miniProject2/internal/account/model/domain"
	"miniProject2/internal/account/model/entity"
	"miniProject2/internal/account/repository"
	"miniProject2/pkg/helper"
	"miniProject2/pkg/security"
	"sync"
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

// UpdateAdminStatusByID implements AccountUseCase.
func (uc *AccountUseCaseImpl) UpdateAdminStatusByID(reqReg domain.AdminReg, reqActor domain.Actor) (int64, error) {
	var wg sync.WaitGroup
	var result int64
	chErr1 := make(chan error, 1)
	chErr2 := make(chan error, 1)
	chInt := make(chan int64, 1)

	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	etAdminReg := entity.AdminReg{
		AdminId: reqReg.AdminId,
		Status:  reqReg.Status,
	}
	etActor := entity.Actor{
		ID:         reqActor.ID,
		IsVerified: reqActor.IsVerified,
		IsActive:   reqActor.IsActive,
	}

	// update admin_reg status only
	// i, err := uc.AccountRepository.UpdateAdminRegStatusByAdminID(tx, etAdminReg)
	// if err != nil {
	// 	return 0, err
	// }

	// update actor is_verified & is_active
	// _, err = uc.AccountRepository.UpdateAdminStatusByAdminID(tx, etActor)
	// if err != nil {
	// 	return 0, err
	// }
	wg.Add(1)
	go func() {
		defer wg.Done()
		i, err := uc.AccountRepository.UpdateAdminRegStatusByAdminID(tx, etAdminReg)
		if err != nil {
			chErr1 <- err
			return
		}
		chInt <- i
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = uc.AccountRepository.UpdateAdminStatusByAdminID(tx, etActor)
		if err != nil {
			chErr2 <- err
			return
		}
	}()
	wg.Wait()

	select {
	case result = <-chInt:
		return result, nil
	case err = <-chErr1:
		return 0, err
	case err = <-chErr2:
		return 0, err
	}
}

// GetAllApprovalAdmin implements AccountUseCase.
func (uc *AccountUseCaseImpl) GetAllApprovalAdmin() ([]domain.AdminReg, error) {
	result := make([]domain.AdminReg, 0)
	tx, err := uc.DB.Begin()
	if err != nil {
		return result, err
	}
	defer helper.CommitOrRollback(err, tx)

	res, err := uc.AccountRepository.GetAllApprovalAdmin(tx)
	if err != nil {
		return result, err
	}

	return DTOListAdminReg(res), nil
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
	res := DTOActor(result)
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
