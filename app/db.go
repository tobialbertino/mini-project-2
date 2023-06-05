package app

import (
	"database/sql"
	"fmt"
	"miniProject2/pkg/config"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(cfg *config.Config) *sql.DB {
	DNS := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Protocol,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	db, err := sql.Open("mysql", DNS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("success to connect database MySQL")
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
