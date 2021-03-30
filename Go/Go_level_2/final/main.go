package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AleksandrMac/GeekBrains/Go/Go_level_2/final/find"
)

var (
	delete *bool
	dir    *string
)

func init() {
	delete = flag.Bool("delete", true, "used to remove duplicate files")
	dir = flag.String("dir", ".", "directory to scan")
	flag.Parse()
}

func main() {
	fmt.Println(find.ReadDir(".."))
	duplicateList, err := find.GetDuplicate(*dir)
	if err != nil {
		log.Println(err)
	}

	for _, listPath := range duplicateList {
		if len(listPath) > 1 {
			fmt.Println("Найдены дублирующиеся файлы:")
			for i, it := range listPath {
				fmt.Printf(" \t%d) %q\n", i, it)
			}
			if *delete {
				find.DeleteDuplicateFiles(listPath)
			}
		}
	}
}
