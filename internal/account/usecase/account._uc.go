package usecase

import "miniProject2/internal/account/model/domain"

type AccountUseCase interface {
	// auth

	// actor
	AddActor(req domain.Actor) (int64, error)
}
