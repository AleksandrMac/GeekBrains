package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sync"
)

var (
	delete *bool
	dir    *string
)

func init() {
	delete = flag.Bool("delete", false, "used to remove duplicate files")
	dir = flag.String("dir", "..", "directory to scan")
	flag.Parse()
}

func main() {
	// if len(flag.Args()) > 0 {
	// 	fmt.Printf("unrecognized argument: %q", flag.Args())
	// 	return
	// }
	list, _ := ReadDirP(*dir)
	fmt.Println(*delete, *dir)
	for _, item := range list {
		fmt.Println(item)
	}
}

// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func ReadDir(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := func() ([]string, error) {
		list, err := f.ReadDir(-1)
		if err != nil {
			return nil, err
		}
		listStr := make([]string, 0, len(list))
		for _, item := range list {
			if item.IsDir() {
				list1, err := ReadDir(dirname + "\\" + item.Name())
				if err != nil {
					return nil, err
				}

				newSlice := make([]string, len(listStr), cap(listStr)+cap(list1))
				copy(newSlice, listStr)
				listStr = newSlice
				listStr = append(listStr, list1...)
			} else {
				listStr = append(listStr, dirname+"\\"+item.Name())
			}
		}
		return listStr, nil
	}()
	f.Close()
	if err != nil {
		return nil, err
	}

	//sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}

// ReadDirP reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func ReadDirP(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	list, err := func() ([]string, error) {
		list, err := f.ReadDir(-1)
		if err != nil {
			return nil, err
		}
		listStr := make(chan string)
		for _, item := range list {
			wg.Add(1)
			go func(item fs.DirEntry) {
				defer wg.Done()
				if item.IsDir() {
					list1, err := ReadDir(dirname + "\\" + item.Name())
					if err != nil {
						log.Println(err)
					}
					for _, item := range list1 {
						listStr <- item
					}
				} else {
					listStr <- dirname + "\\" + item.Name()
				}

			}(item)
		}
		wg.Wait()
		listOut := []string{}
		for val := range listStr {
			listOut = append(listOut, val)
		}
		return listOut, nil
	}()
	f.Close()
	if err != nil {
		return nil, err
	}

	//sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}
