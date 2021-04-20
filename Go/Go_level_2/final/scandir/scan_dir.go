package scandir

import (
	"fmt"
	"hash/adler32"
	"io/fs"
	"os"
	"sync"

	"github.com/spf13/afero"
	"go.uber.org/zap"
)

type ScanDir struct {
	FS  afero.Fs
	Log *zap.Logger
}

func NewScanDir(afs afero.Fs, logger *zap.Logger) *ScanDir {
	dir := new(ScanDir)
	dir.FS = afs
	dir.Log = logger
	return dir
}

// FindDuplicate reads the directory named by dirname and returns
// a duplicate filename.
func (f *ScanDir) FindDuplicate(fileList []string) map[uint32][]string {
	fileSum := make(map[uint32][]string)
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(fileList))
	for _, path := range fileList {
		go func(path string) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()

			sum, err := f.GetSum(path)
			if err != nil {
				f.Log.Error(err.Error(), zap.String("func", "GetSum"))
				return
			}
			fileSum[sum] = append(fileSum[sum], path)
		}(path)
	}
	wg.Wait()

	if len(fileSum) == 0 {
		return nil
	}

	result := make(map[uint32][]string)
	for key, list := range fileSum {
		if len(list) > 1 {
			result[key] = list
		}
	}
	return result
}

// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func (f *ScanDir) ScanDir(dirname string) ([]string, error) {
	list := make([]string, 0, 32)
	FSUTILS := &afero.Afero{Fs: f.FS}
	err := FSUTILS.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			list = append(list, path)
		}
		return err
	})
	return list, err
}

// GetSum returns a hash sum obtained
// from a string consisting of the file name and its size
func (f *ScanDir) GetSum(path string) (uint32, error) {
	fileStat, err := afero.Afero{Fs: f.FS}.Stat(path)
	if err != nil {
		return 0, fmt.Errorf("error \"%w\"\n\toccurred in function \"getSum(%q)\"", err, path)
	}
	return adler32.Checksum([]byte(fileStat.Name() + fmt.Sprintln(fileStat.Size()))), nil
}

// DeleteDuplicateFiles procedure for removing duplicate files
func (f *ScanDir) DeleteDuplicateFiles(listPath []string) (deletedFiles []string, err error) {
	var ind uint16
	for err != nil {
		fmt.Println("Введите целое число >= 0")
		_, err = fmt.Fscan(os.Stdin, &ind)
	}
	err = fmt.Errorf("")
	for i, item := range listPath {
		if uint16(i) != ind {
			if errInside := os.Remove(item); errInside != nil {
				err = fmt.Errorf(", %w%s", errInside, err)
			} else {
				deletedFiles = append(deletedFiles, item)
			}
		}
	}
	return
}
