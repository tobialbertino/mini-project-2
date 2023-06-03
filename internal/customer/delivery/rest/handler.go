package rest

import (
	"miniProject2/internal/customer/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	CustomerUseCase usecase.CustomertUseCase
}

func NewHandler(CustomerUC usecase.CustomertUseCase) *Handler {
	return &Handler{
		CustomerUseCase: CustomerUC,
	}
}

func (h *Handler) Route(app *gin.Engine) {
	// account
	a := NewCustomerHandler(h.CustomerUseCase)
	a.Route(app)
}
