package DButil

import (
	"fmt"
	"testing"
	"urlServer/initconfig"
)

func TestInitDB(t *testing.T) {
	initconfig.FinishInit("../config")
	db := InitDB()
	db.Ping()

	result, err := db.Exec(`insert into person values (null,"tt",11)`)
	if err != nil {
		return
	}
	fmt.Println(result.LastInsertId())
}
