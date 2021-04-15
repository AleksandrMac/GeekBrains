package main

import (
	"flag"
	"fmt"

	"github.com/AleksandrMac/GeekBrains/Go/Go_level_2/final/find"
	"go.uber.org/zap"
)

var (
	delete *bool
	dir    *string
)

func init() {
	delete = flag.Bool("delete", true, "used to remove duplicate files")
	dir = flag.String("dir", "final\\test", "directory to scan")
	flag.Parse()
}

func main() {
	//fmt.Println(find.ReadDir(".."))
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	duplicateList, err := find.GetDuplicate(*dir)
	if err != nil {
		logger.Error(err.Error(),
			zap.String("pkg", "find"),
			zap.String("func", "GetDuplicate"),
		)
	}

	for _, listPath := range duplicateList {
		if len(listPath) > 1 {
			fmt.Println("Найдены дублирующиеся файлы:")
			for i, it := range listPath {
				fmt.Printf(" \t%d) %q\n", i, it)
			}
			if *delete {
				find.DeleteDuplicateFiles(listPath, logger)
			}
		}
	}
}
