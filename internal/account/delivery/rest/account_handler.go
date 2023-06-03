package rest

import (
	"errors"
	"miniProject2/exception"
	"miniProject2/internal/account/model/domain"
	"miniProject2/internal/account/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AccountHandler struct {
	AccountUseCase usecase.AccountUseCase
}

func NewAccountHandler(AccountUC usecase.AccountUseCase) *AccountHandler {
	return &AccountHandler{
		AccountUseCase: AccountUC,
	}
}

// TODO: Implement Authentications through middleware
func (h *AccountHandler) Route(app *gin.Engine) {
	g := app.Group("/account")

	g.POST("", h.AddActor)
	g.POST("/login", h.Login)
}

func (h *AccountHandler) AddActor(c *gin.Context) {
	var req ReqAddActor

	err := c.ShouldBindJSON(&req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			exception.ValidationErrorTranslation(ve, c)
			return
		}
		exception.NewClientError(400, err.Error(), c)
		return
	}

	domain := domain.Actor{
		Username: req.Username,
		Password: req.Password,
	}

	result, err := h.AccountUseCase.AddActor(domain)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := RowsAffected{
		Message:      "Success",
		RowsAffected: result,
	}

	c.JSON(http.StatusOK, res)
}

func (h *AccountHandler) Login(c *gin.Context) {
	var req ReqAddActor

	err := c.ShouldBindJSON(&req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			exception.ValidationErrorTranslation(ve, c)
			return
		}
		exception.NewClientError(http.StatusBadRequest, "Bad Request", c)
		return
	}

	domain := domain.Actor{
		Username: req.Username,
		Password: req.Password,
	}

	result, err := h.AccountUseCase.VerifyActorCredential(domain)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := WebResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    ToResponse(result),
	}

	c.JSON(http.StatusOK, res)
}
