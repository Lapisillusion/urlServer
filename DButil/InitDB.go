package DButil

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"urlServer/initconfig"
)

func InitDB() *sql.DB {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		initconfig.Get("user"),
		initconfig.Get("password"),
		initconfig.Get("dburl"),
		initconfig.Get("dbport"),
		initconfig.Get("dbname"))
	db, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal("init db failed\n", err)
	}
	db.SetMaxOpenConns(32)
	db.SetMaxIdleConns(16)
	return db
}
