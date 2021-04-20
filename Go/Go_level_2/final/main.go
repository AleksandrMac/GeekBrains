package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/AleksandrMac/GeekBrains/Go/Go_level_2/final/scandir"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

func main() {
	var (
		del           *bool
		dir           *string
		fileList      []string
		duplicateList map[uint32][]string
		err           error
	)

	del = flag.Bool("delete", true, "used to remove duplicate files")
	dir = flag.String("dir", "..\\final\\test", "directory to scan")
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer func() {
		err = logger.Sync()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	sd := scandir.NewScanDir(afero.NewOsFs(), logger.With(zap.String("pkg", "scan_dir")))

	fileList, err = sd.ScanDir(*dir)
	if err != nil {
		sd.Log.Error(err.Error())
		return
	}
	duplicateList = sd.FindDuplicate(fileList)

	if duplicateList == nil {
		fmt.Println("Дублирующиеся файлы не найдены.")
		return
	}

	fmt.Println("Найдены дублирующиеся файлы.")
	for _, val := range duplicateList {
		for i, it := range val {
			fmt.Printf(" \t%d) %q\n", i, it)
		}
		if *del {
			list, err := sd.DeleteDuplicateFiles(val)

			loggerDel := logger.With(zap.String("func", "DeleteDuplicateFiles"))
			if err != nil {
				loggerDel.Error(err.Error())
			}
			if len(list) > 0 {
				loggerDel.Info(strings.Join(list, ","))
			}
		}
	}
}
