package initconfig

import (
	"io"
	"log"
	"os"
	"strings"
)

var initConfig = make(map[string]string)

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
		if len(pair) != 2 {
			continue
		}
		k, v := pair[0], pair[1]
		initConfig[k] = v
	}
}

func Get(key string) string {
	v, ok := initConfig[key]
	if !ok {
		log.Fatal("Missing necessary configuration\n")
	}
	return v
}
