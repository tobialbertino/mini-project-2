package rest

import (
	"miniProject2/modules/account/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	AccountUseCase usecase.AccountUseCase
}

func NewHandler(userUC usecase.AccountUseCase) *Handler {
	return &Handler{
		AccountUseCase: userUC,
	}
}

func (h *Handler) Route(app *gin.Engine) {
	// account
	a := NewAccountHandler(h.AccountUseCase)
	a.Route(app)
}
