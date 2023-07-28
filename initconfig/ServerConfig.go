package initconfig

import (
	"io"
	"log"
	"os"
	"strings"
)

var InitConfig = make(map[string]string)

func FinishInit(path string) {
	file, err := os.Open("./" + path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	str := string(bytes)
	entrys := strings.Split(str, "\n")
	for _, e := range entrys {
		e = strings.TrimSpace(e)
		pair := strings.Split(e, ":")
		k, v := pair[0], pair[1]
		InitConfig[k] = v
	}
}
