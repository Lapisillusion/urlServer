package initconfig

import (
	"fmt"
	"testing"
)

func TestFinishInit(t *testing.T) {
	FinishInit("config")
	fmt.Println(InitConfig)
}
