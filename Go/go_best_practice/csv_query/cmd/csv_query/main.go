// main
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	filepath "path"
	"strings"
	"syscall"
	"time"

	"github.com/AleksandrMac/GeekBrains/Go/go_best_practice/csv_query/csv"
	toml "github.com/pelletier/go-toml"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

type Config struct {
	Head    csv.Head      `json:"head" yaml:"head"`
	Sep     string        `json:"sep" yaml:"sep"`
	TimeOut time.Duration `json:"timeOut" yaml:"timeOut"`
	Log     zap.Config    `json:"log" yaml:"log"`
}

var (
	config    Config
	GitCommit string
	query     *string
)

func main() {
	query = flag.String("query", "", "Use example --query=\"continent='Asia' AND date>'2020-04-14'\"")
	flag.Parse()

	fs := afero.NewOsFs()
	buf, err := afero.ReadFile(fs, "configs/config.toml")
	if err != nil {
		log.Default().Println(fmt.Errorf("configuration error: %w", err).Error())
		return
	}
	err = toml.Unmarshal(buf, &config)
	if err != nil {
		log.Default().Println(fmt.Errorf("configuration error: %w", err).Error())
		return
	}
	logger, _ := zap.NewProduction()
	//logger, err := config.Log.Build()
	if err != nil {
		log.Default().Println(err)
	}
	defer logger.Sync()

	//logger.Info(*query)

	path, err := os.Getwd()
	if err != nil {
		logger.Fatal(err.Error())
	}
	fmt.Println("Working directory: ", path)
	fmt.Println("GitCommit: ", GitCommit)
	fmt.Println()

	fmt.Println(filepath.Join(path, config.Head.Path))
	file, err := os.Open(filepath.Join(path, config.Head.Path))
	if err != nil {
		logger.Error(fmt.Errorf("file open error: %w", err).Error())
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	if config.Head.Fields == nil {
		fmt.Printf(
			`В config.toml не найдена информация о названиях полей.
Использовать первую строку в файле %s для инициализации полей?

Нажмите Y(да)/N(нет, завершить)`, config.Head.Path)
		var answer string // = "y"
		fmt.Scan(&answer)
		switch answer {
		case "Y", "y":
			var str []byte
			str, _, err = reader.ReadLine()
			config.Head.Fields = csv.GetFields(string(str), config.Sep)
		default:
			return
		}
	}

	//fmt.Println(*query)
	ctx, cancel := context.WithTimeout(context.Background(), config.TimeOut*time.Second)
	go watchSignals(cancel, logger)
	defer cancel()

	for err == nil {
		var str []byte
		str, _, err = reader.ReadLine()
		if len(str) > 0 {
			time.Sleep(100 * time.Millisecond)
			go func(ctx context.Context) {
				select {
				case <-ctx.Done():
					return
				default:
					row := config.Head.NewRow()
					row.Values = strings.Split(string(str), config.Sep)
					if row.IsMatch(*query) {
						fmt.Println(row.Values)
					}
				}
			}(ctx)
		}
		select {
		case <-ctx.Done():
			logger.Error(ctx.Err().Error())
			return
		default:
			continue
		}
	}
}

func watchSignals(cancel context.CancelFunc, logger *zap.Logger) {
	osSignalChan := make(chan os.Signal, 1)

	signal.Notify(osSignalChan,
		syscall.SIGINT)
	sig := <-osSignalChan
	logger.Error(fmt.Sprintf("got signal %q", sig.String()))

	// если сигнал получен, отменяем контекст работы
	cancel()
}

// export GIT_COMMIT=$(git rev-list -1 HEAD) && \ go build -ldflags "-X main.GitCommit=$GIT_COMMIT"
