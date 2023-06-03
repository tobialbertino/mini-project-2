package app

import (
	"database/sql"
	accHandler "miniProject2/internal/account/delivery/rest"
	accRepo "miniProject2/internal/account/repository"
	accUC "miniProject2/internal/account/usecase"

	customerHandler "miniProject2/internal/customer/delivery/rest"
	customerRepo "miniProject2/internal/customer/repository"
	customerUC "miniProject2/internal/customer/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine, DB *sql.DB) {
	// account setup
	accountRepo := accRepo.NewAccountRepository()
	accountUC := accUC.NewAccountUseCase(accountRepo, DB)
	AccountHandler := accHandler.NewHandler(accountUC)
	AccountHandler.Route(app)

	// Customer setup
	customerRepo := customerRepo.NewCustomerRepository()
	customerUC := customerUC.NewCustomerUseCase(customerRepo, DB)
	CustomerHandler := customerHandler.NewHandler(customerUC)
	CustomerHandler.Route(app)
}
