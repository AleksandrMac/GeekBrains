package config

import (
	"bufio"
	"os"
	"strings"
)

// Config -
func Config(file *os.File) (conf map[string]string, err map[string]error) {
	//fmt.Println(filename)
	//return nil, nil

	scanner := bufio.NewScanner(file)
	conf = make(map[string]string)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ": ")
		conf[arr[0]] = arr[1]
	}
	return
}
