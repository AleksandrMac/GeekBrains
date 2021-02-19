package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
)

// Config - возвращает данные из конфигурационного файла
//первый аргумент имя файла,
//второй тип структуры файла, есл не указан определяется из разрешения файла
//далее перечень параметров необходимых для считывания, в случае отсутствия, возвращается полный список
func Config(params ...string /*    file *os.File, format string = "auto"*/) (conf map[string]interface{}, err map[string]error) {

	conf = make(map[string]interface{})
	err = make(map[string]error)
	lenParams := uint16(len(params))

	if len(params[0]) < 3 && lenParams > 0 {
		err["file"] = errors.New("ожидается ввод имени файла")
		return
	}

	var file *os.File
	file, err["file"] = os.Open(params[0])
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

	//определяем какую структуру данных необходимо обработать
	var fileStruct string
	if lenParams > 1 {
		fileStruct = params[1]
	} else {
		fileStruct = dataStructure(params[0])
	}
	//обрабатываем данные
	switch fileStruct {
	case "yaml":
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			arr := strings.Split(scanner.Text(), ": ")
			conf[arr[0]] = arr[1]
		}
		//если существуют на входе определенные параметры, то выводим их
		if lenParams > 2 {
			temp := make(map[string]interface{})
			for i := uint16(2); i < lenParams; i++ {
				temp[params[i]] = conf[params[i]]
			}
			conf = temp
		}
		return
	case "json":
		jsonStream := ""
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jsonStream += scanner.Text()
		}
		if err := json.Unmarshal([]byte(jsonStream), &conf); err != nil {
			panic(err)
		}
		//если существуют на входе определенные параметры, то выводим их
		if lenParams > 2 {
			temp := make(map[string]interface{})
			for i := uint16(2); i < lenParams; i++ {
				temp[params[i]] = conf[params[i]]
			}
			conf = temp
		}

		return
	default:
		err["struct"] = errors.New("<" + fileStruct + "> Структура данных не известна, попробуйте вторым параметром указать: yaml, json")
		return
	}

}

func dataStructure(fileName string) string {
	slice := strings.Split(fileName, ".")
	return slice[len(slice)-1]
}
