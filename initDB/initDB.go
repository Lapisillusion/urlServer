package initDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"urlServer/initconfig"
)

func InitDB() *sql.DB {
	sourse := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		initconfig.GetConfig("user"),
		initconfig.GetConfig("password"),
		initconfig.GetConfig("url"),
		initconfig.GetConfig("port"),
		initconfig.GetConfig("dbname"))
	db, err := sql.Open("mysql", sourse)
	if err != nil {
		log.Fatal("init db failed\n", err)
	}
	db.SetMaxOpenConns(32)
	db.SetMaxIdleConns(16)
	return db
}
