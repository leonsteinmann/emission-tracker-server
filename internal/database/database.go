package database

import (
	"database/sql"
	"fmt"

	"emission-tracker-server/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(cfg config.DBConfig) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
