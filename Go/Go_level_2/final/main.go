package main

import (
	"flag"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"log"
	"os"
)

var (
	delete  *bool
	dir     *string
	fileSum map[uint32][]string
)

func init() {
	delete = flag.Bool("delete", true, "used to remove duplicate files")
	dir = flag.String("dir", ".", "directory to scan")
	flag.Parse()
}

func main() {
	fileSum = make(map[uint32][]string)
	// if len(flag.Args()) > 0 {
	// 	fmt.Printf("unrecognized argument: %q", flag.Args())
	// 	return
	// }
	list, err := ReadDir(*dir)
	if err != nil {
		log.Println(err)
		return
	}
	crc := crc32.NewIEEE()
	for _, item := range list {
		file, _ := os.Open(item)
		fileStat, _ := file.Stat()
		fmt.Println(item, fileStat.Size())
		fmt.Println([]byte(fileStat.Name() + fmt.Sprintln(fileStat.Size())))
		crc.Write([]byte(fileStat.Name() + fmt.Sprintln(fileStat.Size())))
		//sum := crc.Sum32()
		sum := adler32.Checksum([]byte(fileStat.Name() + fmt.Sprintln(fileStat.Size())))
		if fileSum[sum] == nil {
			fileSum[sum] = make([]string, 0, 4)
		}
		fmt.Println(sum)
		fileSum[sum] = append(fileSum[sum], item)
	}

	for _, items := range fileSum {
		if len(items) > 1 {
			fmt.Println("Найдены дублирующиеся файлы:")
			for i, it := range items {
				fmt.Printf(" \t%d) %q\n", i, it)
			}
			if *delete {
				fmt.Println("Введите индекс файла, который необходимо сохранить:")
				var ind uint16
				err := fmt.Errorf("")
				for err != nil {
					fmt.Println("Введите целое число >= 0")
					_, err = fmt.Fscan(os.Stdin, &ind)
				}
				for i, item := range items {
					if uint16(i) != ind {
						fmt.Println("Удаляется файл: ", item)
						fmt.Println(os.Remove(item))
					}
				}
			}
		}
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
