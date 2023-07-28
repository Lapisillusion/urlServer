package initDB

import (
	"testing"
	"urlServer/initconfig"
)

func TestInitDB(t *testing.T) {
	initconfig.FinishInit("../config")
	db := InitDB()
	db.Ping()
}
