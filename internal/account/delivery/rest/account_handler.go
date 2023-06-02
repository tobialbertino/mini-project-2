package rest

import (
	"miniProject2/internal/account/model/domain"
	"miniProject2/internal/account/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	AccountUseCase usecase.AccountUseCase
}

func NewAccountHandler(AccountUC usecase.AccountUseCase) *AccountHandler {
	return &AccountHandler{
		AccountUseCase: AccountUC,
	}
}

func (h *AccountHandler) Route(app *gin.Engine) {
	g := app.Group("/account")

	g.POST("", h.AddActor)
}

func (h *AccountHandler) AddActor(c *gin.Context) {
	var req ReqAddActor

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	domain := domain.Actor{
		Username: req.Username,
		Password: req.Password,
	}

	result, err := h.AccountUseCase.AddActor(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := RowsAffected{
		Message:      "Success",
		RowsAffected: result,
	}

	c.JSON(http.StatusOK, res)
}
