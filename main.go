package main

import (
	"log"
	"miniProject2/app"
	"miniProject2/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// load config
	config, err := config.LoadConfig()
	if err != nil {
		log.Printf("error loading config: %s", err)
	}

	// database
	DB := app.NewDB(config)

	// init router
	r := gin.Default()
	app.InitRouter(r, DB)

	r.Run(config.Server.Port)
}
