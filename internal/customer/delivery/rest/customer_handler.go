package rest

import (
	"errors"
	"miniProject2/exception"
	"miniProject2/internal/customer/model/domain"
	"miniProject2/internal/customer/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerHandler struct {
	CustomerUseCase usecase.CustomertUseCase
}

func NewCustomerHandler(CustomerUC usecase.CustomertUseCase) *CustomerHandler {
	return &CustomerHandler{
		CustomerUseCase: CustomerUC,
	}
}

// TODO: Implement Authentications through middleware
func (h *CustomerHandler) Route(app *gin.Engine) {
	g := app.Group("/customer")

	g.GET("", h.GetAllCustomer)
}

func (h *CustomerHandler) GetAllCustomer(c *gin.Context) {
	var req ReqGetAllCustomer
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

	dm := domain.Customer{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	result, err := h.CustomerUseCase.GetAllCustomer(dm)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := WebResponse{
		Message: "Success",
		Data:    ToResponseListCustomer(result),
	}

	c.JSON(http.StatusOK, res)
}
