package usecase

import "miniProject2/internal/account/model/domain"

type AccountUseCase interface {
	// auth
	VerifyActorCredential(req domain.Actor) (domain.Actor, error)
	// actor
	AddActor(req domain.Actor) (int64, error)
}
