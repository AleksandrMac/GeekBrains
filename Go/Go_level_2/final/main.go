package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/AleksandrMac/GeekBrains/Go/Go_level_2/final/scan_dir"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

var (
	delete *bool
	dir    *string
)

func init() {
	delete = flag.Bool("delete", true, "used to remove duplicate files")
	dir = flag.String("dir", "..\\final\\test", "directory to scan")
	flag.Parse()
}

func main() {
	var (
		fileList      []string
		duplicateList map[uint32][]string
		err           error
	)
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sd := scan_dir.NewScanDir(afero.NewOsFs(), logger.With(zap.String("pkg", "scan_dir")))

	fileList, err = sd.ScanDir(*dir)
	if err != nil {
		sd.Log.Error(err.Error())
		os.Exit(1)
	}
	duplicateList = sd.FindDuplicate(fileList)

	if duplicateList == nil {
		fmt.Println("Дублирующиеся файлы не найдены.")
		os.Exit(0)
	}

	fmt.Println("Найдены дублирующиеся файлы.")
	for key, _ := range duplicateList {
		for i, it := range duplicateList[key] {
			fmt.Printf(" \t%d) %q\n", i, it)
		}
		if *delete {
			list, err := sd.DeleteDuplicateFiles(duplicateList[key])

			loggerDel := logger.With(zap.String("func", "DeleteDuplicateFiles"))
			if err != nil {
				loggerDel.Error(err.Error())
			}
			if len(list) > 0 {
				loggerDel.Info(strings.Join(list, ","))
			}
		}
	}

	// 	}
	// }
}
