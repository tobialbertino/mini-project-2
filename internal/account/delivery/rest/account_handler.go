package rest

import (
	"miniProject2/exception"
	"miniProject2/internal/account/model/domain"
	"miniProject2/internal/account/usecase"
	"miniProject2/pkg/middleware"
	"net/http"
	"strconv"

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
	app.POST("/login", h.Login) // Generate token JWT

	g := app.Group("/account", middleware.Auth()) // using middleware
	g.GET("", h.GetAllAdmin)                      // TODO: implement goroutine, sometimes errors bad connection, and a busy buffer
	g.POST("", h.AddActor)

	// only super_admin
	g.GET("/admin-reg", middleware.AuthSuperAdmin(), h.GetAllAppovalAdmin)
	g.PUT("/admin-reg", middleware.AuthSuperAdmin(), h.UpdateAdminStatus)
	g.DELETE("/admin-reg", middleware.AuthSuperAdmin(), h.DeleteAdminByID)
}

func (h *AccountHandler) GetAllAdmin(c *gin.Context) {
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		exception.NewClientError(400, err.Error(), c)
		return
	}
	username := c.Query("username")

	dm := domain.Actor{
		Username: username,
	}
	dmPaging := domain.Pagination{
		Page: pageInt,
	}

	result, err := h.AccountUseCase.GetAllAdmin(dm, dmPaging)
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	// combine result
	combineResult := ToResGetAllAdminWithPaging(result)

	res := WebResponse{
		Message: "Success",
		Data:    combineResult,
	}

	c.JSON(http.StatusOK, res)
}

func (h *AccountHandler) DeleteAdminByID(c *gin.Context) {
	var req ReqIDActor
	err := c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	dm := domain.Actor{
		ID: req.ID,
	}

	result, err := h.AccountUseCase.DeleteAdminByID(dm)
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

func (h *AccountHandler) UpdateAdminStatus(c *gin.Context) {
	var req ReqUpdateAdminStatus

	err := c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
		return
	}

	dmActor := domain.Actor{
		ID:         req.AdminID,
		IsVerified: req.IsVerified,
		IsActive:   req.IsActive,
	}
	dmAdminReg := domain.AdminReg{
		AdminId: req.AdminID,
		Status:  req.Status,
	}

	result, err := h.AccountUseCase.UpdateAdminStatusByID(dmAdminReg, dmActor)
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

func (h *AccountHandler) GetAllAppovalAdmin(c *gin.Context) {

	result, err := h.AccountUseCase.GetAllApprovalAdmin()
	if err != nil {
		exception.NewInternalError(http.StatusInternalServerError, err.Error(), c)
		return
	}

	res := WebResponse{
		Message: "Success",
		Data:    ResponseListAdminReg(result),
	}

	c.JSON(http.StatusOK, res)
}

func (h *AccountHandler) AddActor(c *gin.Context) {
	var req ReqAddActor

	err := c.ShouldBindJSON(&req)
	if err != nil {
		exception.BindJSONError(err, c)
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
		exception.BindJSONError(err, c)
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
		Data:    result,
	}

	c.JSON(http.StatusOK, res)
}
