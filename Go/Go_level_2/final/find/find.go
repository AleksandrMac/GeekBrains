package find

import (
	"fmt"
	"hash/adler32"
	"os"
	"sync"

	"go.uber.org/zap"
)

// GetDuplicate reads the directory named by dirname and returns
// a duplicate filename.
func GetDuplicate(dirname string) ([][]string, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger = logger.With(zap.String("pkg", "find"))

	list, err := ReadDir(dirname, logger.With(zap.String("func", "ReadDir")))
	if err != nil {
		logger.Error(err.Error(), zap.String("func", "ReadDir"))
		return nil, err
	}

	fileSum := make(map[uint32][]string)
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(list))
	for _, path := range list {
		go func(path string) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()

			sum, err := GetSum(path)
			if err != nil {
				logger.Error(err.Error(), zap.String("func", "GetSum"))
				return
			}
			fileSum[sum] = append(fileSum[sum], path)
		}(path)
	}
	wg.Wait()

	result := make([][]string, 0, 100)
	for _, list := range fileSum {
		if len(list) > 1 {
			result = append(result, list)
		}
	}
	len := len(result)
	return result[:len], nil
}

// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func ReadDir(dirname string, logger *zap.Logger) ([]string, error) {
	logger.Info("Scan start: " + dirname)

	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	list, err := func() ([]string, error) {
		list, err := f.ReadDir(-1)
		if err != nil {
			return nil, err
		}
		listStr := make([]string, 0, len(list))
		for _, item := range list {
			if item.IsDir() {
				//вымышленная паника
				panic(dirname + "\\" + item.Name())
				list1, err := ReadDir(dirname+"\\"+item.Name(), logger)
				if err == nil {
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
	if err != nil {
		return nil, err
	}
	logger.Info("Scan end: "+dirname, zap.String("func", "ReadDir"))
	return list, nil
}

// GetSum returns a hash sum obtained
// from a string consisting of the file name and its size
func GetSum(path string) (uint32, error) {
	fileStat, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("error \"%w\"\n\toccurred in function \"getSum(%q)\"", err, path)
	}
	return adler32.Checksum([]byte(fileStat.Name() + fmt.Sprintln(fileStat.Size()))), nil
}

// DeleteDuplicateFiles procedure for removing duplicate files
func DeleteDuplicateFiles(listPath []string, logger *zap.Logger) {
	fmt.Println("Введите индекс файла, который необходимо сохранить:")
	var ind uint16
	err := fmt.Errorf("")
	for err != nil {
		fmt.Println("Введите целое число >= 0")
		_, err = fmt.Fscan(os.Stdin, &ind)
	}
	for i, item := range listPath {
		if uint16(i) != ind {
			err = os.Remove(item)
			if err != nil {
				logger.Error(err.Error(),
					zap.String("pkg", "find"),
					zap.String("func", "DeletedDuplicatesFiles"),
					zap.Bool("stat", false),
				)
			} else {
				logger.Info(item,
					zap.String("pkg", "find"),
					zap.String("func", "DeletedDuplicatesFiles"),
					zap.Bool("stat", true),
				)
			}
		}
	}
}
