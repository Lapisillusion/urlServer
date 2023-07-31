package jsonutil

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	f, _ := os.Open("./test.json")
	bytes, _ := io.ReadAll(f)
	p, _ := Parser(bytes)
	fmt.Println(p)
}
