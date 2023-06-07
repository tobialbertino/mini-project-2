package usecase

import (
	"database/sql"
	"errors"
	"miniProject2/internal/account/model/domain"
	"miniProject2/internal/account/model/entity"
	"miniProject2/internal/account/repository"
	"miniProject2/pkg/helper"
	"miniProject2/pkg/security"
	"miniProject2/pkg/tokenize"
	"sync"
	"time"
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

// GetAllAdmin implements AccountUseCase.
func (uc *AccountUseCaseImpl) GetAllAdmin(req domain.Actor, pagi domain.Pagination) (domain.ListActorWithPaging, error) {
	var (
		err       error
		wg        sync.WaitGroup
		resPaging entity.Pagination
		result    []entity.Actor
	)

	chListAdmin := make(chan []entity.Actor, 1)
	chPaging := make(chan entity.Pagination, 1)
	errListAdmin := make(chan error, 1)
	errPagination := make(chan error, 1)

	// ?: Error tx with go routine, temporary solution using db queries, maybe tx MySQL doesn't support query select rows on goroutines
	// tx, err := uc.DB.Begin()
	// if err != nil {
	// 	return domain.ListActorWithPaging{}, err
	// }
	// defer helper.CommitOrRollback(err, tx)

	// define pagination
	etPaging := entity.Pagination{
		Page:       pagi.Page,
		PerPage:    6,                   // always fix 6 data == LIMIT
		Total:      0,                   // after query
		TotalPages: 0,                   // after query, total / PerPage
		Offset:     (pagi.Page - 1) * 6, // (Page-1) * PerPage
	}
	et := entity.Actor{
		Username: req.Username,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		result, err := uc.AccountRepository.GetAllAdmin(uc.DB, et, etPaging)
		if err != nil {
			errListAdmin <- err
		}
		chListAdmin <- result
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Get Total Data
		resPaging, err := uc.AccountRepository.Pagination(uc.DB, etPaging)
		if err != nil {
			errPagination <- err
		}
		chPaging <- resPaging
	}()
	wg.Wait()

	for i := 0; i < 2; i++ {
		select {
		case result = <-chListAdmin:
			continue
		case resPaging = <-chPaging:
			continue
		case err = <-errListAdmin:
			return domain.ListActorWithPaging{}, err
		case err = <-errPagination:
			return domain.ListActorWithPaging{}, err
		}
	}

	totalPages := resPaging.Total / 6
	if resPaging.Total%6 != 0 {
		totalPages++
	}
	etPaging.Total = resPaging.Total
	etPaging.TotalPages = totalPages

	combineRes := domain.ListActorWithPaging{
		Pagination: domain.Pagination(etPaging),
		Admins:     DTOActorList(result),
	}

	return combineRes, nil
}

// DeleteAdminByID implements AccountUseCase.
func (uc *AccountUseCaseImpl) DeleteAdminByID(req domain.Actor) (int64, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, tx)

	entity := entity.Actor{
		ID: req.ID,
	}
	result, err := uc.AccountRepository.DeleteAdminByID(tx, entity)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// UpdateAdminStatusByID implements AccountUseCase.
func (uc *AccountUseCaseImpl) UpdateAdminStatusByID(reqReg domain.AdminReg, reqActor domain.Actor) (int64, error) {
	var (
		wg     sync.WaitGroup
		result int64
	)
	chErr1 := make(chan error, 1)
	chErr2 := make(chan error, 1)
	chInt := make(chan int64, 1)
	chInt2 := make(chan int64, 1)

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

	wg.Add(1)
	go func() {
		defer wg.Done()
		// update admin_reg status only
		i, err := uc.AccountRepository.UpdateAdminRegStatusByAdminID(tx, etAdminReg)
		if err != nil {
			chErr1 <- err
		}
		chInt <- i
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// update actor is_verified & is_active
		i2, err := uc.AccountRepository.UpdateAdminStatusByAdminID(tx, etActor)
		if err != nil {
			chErr2 <- err
		}
		chInt2 <- i2
	}()
	wg.Wait()

	// get 2 channel data
	var totalRowsAffected int64
	for i := 0; i < 2; i++ {
		select {
		case result = <-chInt2:
			totalRowsAffected += result
		case result = <-chInt:
			totalRowsAffected += result
		case err = <-chErr1:
			return 0, err
		case err = <-chErr2:
			return 0, err
		}
	}
	return totalRowsAffected, nil
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
func (uc *AccountUseCaseImpl) VerifyActorCredential(req domain.Actor) (domain.ResToken, error) {
	tx, err := uc.DB.Begin()
	if err != nil {
		return domain.ResToken{}, err
	}
	defer helper.CommitOrRollback(err, tx)

	entity := entity.Actor{
		Username: req.Username,
		Password: req.Password,
	}
	result, err := uc.AccountRepository.VerifyActorCredential(tx, entity)
	if err != nil {
		return domain.ResToken{}, err
	}

	// compare password
	isValid := security.CheckPasswordHash(req.Password, result.Password)
	if !isValid {
		return domain.ResToken{}, errors.New("invalid username or password")
	}
	userDetail := DTOActor(result)

	// generate token jwt
	// Create the Claims
	myClaims := tokenize.AccountClaims{
		IDNum:      userDetail.ID,
		RoleID:     userDetail.RoleID,
		IsVerified: userDetail.IsVerified,
		IsActive:   userDetail.IsActive,
		ExpiresAt:  time.Now().Add(time.Hour * 1).Unix(),
	}

	token, err := tokenize.GenerateAccessToken(myClaims)
	if err != nil {
		return domain.ResToken{}, err
	}

	res := domain.ResToken{
		AccessToken: token,
	}
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
