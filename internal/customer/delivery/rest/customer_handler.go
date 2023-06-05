package rest

import (
	"miniProject2/exception"
	"miniProject2/internal/customer/model/domain"
	"miniProject2/internal/customer/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	g.GET("", h.GetAllCustomer) // TODO: implement goroutine
	g.GET("/:id", h.GetCustomerByID)
	g.POST("", h.CreateCustomer)
	g.PUT("/:id", h.UpdateCustomerByID)
	g.DELETE("/:id", h.DeleteCustomerByID)
}

func (h *CustomerHandler) GetAllCustomer(c *gin.Context) {
	var req ReqGetAllCustomer
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		exception.NewClientError(400, err.Error(), c)
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	dm := domain.Customer{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
	dmPaging := domain.Pagination{
		Page: pageInt,
	}

	result, err := h.CustomerUseCase.GetAllCustomer(dm, dmPaging)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	combineResult := ToResGetAllCustomerWithPaging(result)

	res := WebResponse{
		Message: "Success",
		Data:    combineResult,
	}

	c.JSON(http.StatusOK, res)
}

func (h *CustomerHandler) GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		exception.NewClientError(400, err.Error(), c)
	}

	dm := domain.Customer{
		ID: int64(idInt),
	}
	result, err := h.CustomerUseCase.GetCustomerByID(dm)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := WebResponse{
		Message: "Success",
		Data:    ToResponseCustomer(result),
	}

	c.JSON(http.StatusOK, res)
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var req ReqAddCustomer

	err := c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	dm := domain.Customer{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Avatar:    req.Avatar,
	}
	result, err := h.CustomerUseCase.CreateCustomer(dm)
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

func (h *CustomerHandler) UpdateCustomerByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		exception.NewClientError(400, err.Error(), c)
	}

	var req ReqAddCustomer

	err = c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	dm := domain.Customer{
		ID:        int64(idInt),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Avatar:    req.Avatar,
	}
	result, err := h.CustomerUseCase.UpdateCustomerByID(dm)
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

func (h *CustomerHandler) DeleteCustomerByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		exception.NewClientError(400, err.Error(), c)
	}

	dm := domain.Customer{
		ID: int64(idInt),
	}
	result, err := h.CustomerUseCase.DeleteCustomerByID(dm)
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
