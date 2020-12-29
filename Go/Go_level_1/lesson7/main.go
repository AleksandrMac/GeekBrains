package main

import (
	"fmt"
	"lesson7/config"
	"log"
	"os"
)

func main() {
	conf := make(map[string]string)
	err := make(map[string]error)
	var file *os.File
	println(os.Getwd())
	file, err["file"] = os.Open("conf.yaml")
	fmt.Println(file.Name())
	if err["file"] != nil {
		log.Fatal(err)
	}
	delete(err, "file")

	defer func() {
		file.Close()
		if len(err) > 0 {
			log.Fatal(err)
		}
	}()

	conf, err = config.Config(file)

	fmt.Println(conf)
	if len(err) > 0 {
		log.Fatal(err)
	}
}
