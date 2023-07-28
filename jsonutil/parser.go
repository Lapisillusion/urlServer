package jsonutil

import (
	"encoding/json"
	"urlServer/bean"
)

func Parser(sourse []byte) (*bean.Pack, error) {
	var p bean.Pack
	err := json.Unmarshal(sourse, &p)
	return &p, err
}
