package app

import (
	"database/sql"
	accHandler "miniProject2/modules/account/delivery/rest"
	accRepo "miniProject2/modules/account/repository"
	accUC "miniProject2/modules/account/usecase"

	customerHandler "miniProject2/modules/customer/delivery/rest"
	customerRepo "miniProject2/modules/customer/repository"
	customerUC "miniProject2/modules/customer/usecase"

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
