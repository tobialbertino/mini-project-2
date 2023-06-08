package usecase

import (
	"miniProject2/modules/account/model/domain"
)

type AccountUseCase interface {
	// auth
	VerifyActorCredential(req domain.Actor) (domain.ResToken, error) // generate token jwt
	// actor
	AddActor(req domain.Actor) (int64, error)
	// GetAllAdmin With Pagination
	GetAllAdmin(req domain.Actor, pagi domain.Pagination) (domain.ListActorWithPaging, error)

	// super_admin
	GetAllApprovalAdmin() ([]domain.AdminReg, error)
	UpdateAdminStatusByID(reqReg domain.AdminReg, reqActor domain.Actor) (int64, error)
	DeleteAdminByID(req domain.Actor) (int64, error)
}
