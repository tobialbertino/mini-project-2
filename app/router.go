package app

import (
	"database/sql"
	"miniProject2/internal/account/delivery/rest"
	"miniProject2/internal/account/repository"
	"miniProject2/internal/account/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine, DB *sql.DB) {
	// account setup
	accountRepo := repository.NewAccountRepository()
	accountUC := usecase.NewAccountUseCase(accountRepo, DB)
	AccountHandler := rest.NewHandler(accountUC)
	AccountHandler.Route(app)
}
