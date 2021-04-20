// main
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	filepath "path"
	"strings"

	"github.com/AleksandrMac/GeekBrains/Go/go_best_practice/csv_query/csv"
	toml "github.com/pelletier/go-toml"
	"github.com/spf13/afero"
)

type Config struct {
	Head csv.Head
	Sep  string
}

var (
	config    Config
	GitCommit string
)

func main() {
	fs := afero.NewOsFs()
	buf, err := afero.ReadFile(fs, "config.toml")
	if err != nil {
		fmt.Println(fmt.Errorf("configuration error: %w", err))
		return
	}
	err = toml.Unmarshal(buf, &config)
	if err != nil {
		fmt.Println(fmt.Errorf("configuration error: %w", err))
		return
	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Working directory: ", path)
	fmt.Println("GitCommit: ", GitCommit)
	fmt.Println()

	fmt.Println(filepath.Join(path, config.Head.Path))
	file, err := os.Open(filepath.Join(path, config.Head.Path))
	if err != nil {
		fmt.Println(fmt.Errorf("file open error: %w", err))
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	if config.Head.Fields == nil {
		fmt.Printf(
			`В config.toml не найдена информация о названиях полей.\n
Использовать первую строку в файле %s для инициализации полей?\n
Нажмите Y(да)/N(нет, завершить)\n`, config.Head.Path)
		var answer string
		fmt.Scan(&answer)
		switch answer {
		case "Y", "y":

			var str []byte

			str, _, err = reader.ReadLine()
			// str := scanner.Text()
			config.Head.Fields = csv.GetFields(string(str), config.Sep)
		default:
			return
		}
	}

	for err == nil {
		var str []byte
		str, _, err = reader.ReadLine()
		row := config.Head.NewRow()
		row.Values = strings.Split(string(str), config.Sep)
		if row.IsMatch("") {
			fmt.Println(row.Values)
		}
	}
}

// export GIT_COMMIT=$(git rev-list -1 HEAD) && \ go build -ldflags "-X main.GitCommit=$GIT_COMMIT"
