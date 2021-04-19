package main

import (
	"fmt"
	"log"
	"os"

	toml "github.com/pelletier/go-toml"
	"github.com/spf13/afero"
)

type Data struct {
	Path   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
}

type Config struct {
	Data Data
}

var (
	config    Config
	GitCommit string
)

func init() {
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
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Working directory: ", path)
	fmt.Println("GitCommit: ", GitCommit)
	fmt.Println()

}

// export GIT_COMMIT=$(git rev-list -1 HEAD) && \  go build -ldflags "-X main.GitCommit=$GIT_COMMIT"
