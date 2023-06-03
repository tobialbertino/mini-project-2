package usecase

import (
	"miniProject2/internal/account/model/domain"
)

type AccountUseCase interface {
	// auth
	VerifyActorCredential(req domain.Actor) (domain.Actor, error)
	// actor
	AddActor(req domain.Actor) (int64, error)
	GetAllAdmin(req domain.Actor) ([]domain.Actor, error)

	// super_admin
	GetAllApprovalAdmin() ([]domain.AdminReg, error)
	UpdateAdminStatusByID(reqReg domain.AdminReg, reqActor domain.Actor) (int64, error)
	DeleteAdminByID(req domain.Actor) (int64, error)
}
