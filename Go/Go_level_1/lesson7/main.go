package main

import (
	"errors"
	"fmt"
	"lesson7/config"
	"lesson7/valid"
	"log"
)

func main() {
	conf := make(map[string]interface{})
	err := make(map[string]error)

	//conf, err = config.Config("conf.json", "json", "db_url")
	conf, err = config.Config("conf.yaml", "yaml", "db_url")
	//Валидация
	_, ok := conf["port"]
	if ok && !valid.IsPort(conf["port"].(float64)) {
		err["port"] = errors.New("Номер порта указан не корректно")
	}

	_, ok = conf["db_url"]
	if ok {
		_, err["db_url"] = valid.IsURL(conf["db_url"].(string))
	}

	_, ok = conf["jaeger_url"]
	if ok {
		_, err["jaeger_url"] = valid.IsURL(conf["jaeger_url"].(string))
	}

	_, ok = conf["sentry_url"]
	if ok {
		_, err["sentry_url"] = valid.IsURL(conf["sentry_url"].(string))
	}

	fmt.Println(conf)
	if len(err) > 0 {
		log.Fatal(err)
	}
}
