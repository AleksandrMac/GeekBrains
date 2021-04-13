package scan_dir

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type FindTestSuite struct {
	suite.Suite
	sd *ScanDir
	//FS afero.Fs
}

func (fs *FindTestSuite) SetupTest() {
	f := afero.NewMemMapFs()

	f.Mkdir("test", 0644)
	f.Mkdir("test/test_data", 0644)
	if file, err := f.Create("test/file1.txt"); err != nil {
		fs.Errorf(err, "Created file error: %v")
	} else {
		if _, err = file.Write([]byte("test file1")); err != nil {
			fs.Errorf(err, "Writed to file error: %v")
		}
	}

	f.Mkdir("test/test_data/testDir", 0644)
	if file, err := f.Create("test/test_data/file2.txt"); err != nil {
		fs.Errorf(err, "Created file error: %v")
	} else {
		if _, err = file.Write([]byte("test file2")); err != nil {
			fs.Errorf(err, "Writed to file error: %v")
		}
	}

	f.Mkdir("test/test_data/testDir/dir1", 0644)
	f.Mkdir("test/test_data/testDir/dir2", 0644)
	if file, err := f.Create("test/test_data/testDir/file2.txt"); err != nil {
		fs.Errorf(err, "Created file error: %v")
	} else {
		if _, err = file.Write([]byte("test file2")); err != nil {
			fs.Errorf(err, "Writed to file error: %v")
		}
	}

	f.Mkdir("test/test_data/testDir/dir1/dir3", 0644)
	if file, err := f.Create("test/test_data/testDir/dir1/file1.txt"); err != nil {
		fs.Errorf(err, "Created file error: %v")
	} else {
		if _, err = file.Write([]byte("test file1")); err != nil {
			fs.Errorf(err, "Writed to file error: %v")
		}
	}
	if file, err := f.Create("test/test_data/testDir/dir1/file2.txt"); err != nil {
		fs.Errorf(err, "Created file error: %v")
	} else {
		if _, err = file.Write([]byte("test file2")); err != nil {
			fs.Errorf(err, "Writed to file error: %v")
		}
	}

	if file, err := f.Create("test/test_data/testDir/dir2/file1.txt"); err != nil {
		fs.Errorf(err, "Created file error: %v")
	} else {
		if _, err = file.Write([]byte("test file1")); err != nil {
			fs.Errorf(err, "Writed to file error: %v")
		}
	}
	if file, err := f.Create("test/test_data/testDir/dir2/file2.txt"); err != nil {
		fs.Errorf(err, "Created file error: %v")
	} else {
		if _, err = file.Write([]byte("test file2")); err != nil {
			fs.Errorf(err, "Writed to file error: %v")
		}
	}
	logger, _ := zap.NewProduction()
	fs.sd = NewScanDir(f, logger.With(zap.String("pkg", "test_scan")))
}

func (fs *FindTestSuite) TestScanDir() {
	want := []string{
		"test\\file1.txt",
		"test\\test_data\\file2.txt",
		"test\\test_data\\testDir\\dir1\\file1.txt",
		"test\\test_data\\testDir\\dir1\\file2.txt",
		"test\\test_data\\testDir\\dir2\\file1.txt",
		"test\\test_data\\testDir\\dir2\\file2.txt",
		"test\\test_data\\testDir\\file2.txt",
	}
	list, err := fs.sd.ScanDir(".")
	fs.Require().NoError(err)
	fs.Assert().Equal(want, list)
}

func (fs *FindTestSuite) TestFindDuplicate() {
	var (
		fileList      []string
		duplicateList map[uint32][]string
		err           error
	)

	want := map[uint32][]string{
		465765323: {
			"test\\test_data\\testDir\\dir1\\file1.txt",
			"test\\test_data\\testDir\\dir2\\file1.txt",
			"test\\file1.txt",
		},
		466289612: {
			"test\\test_data\\testDir\\file2.txt",
			"test\\test_data\\testDir\\dir1\\file2.txt",
			"test\\test_data\\testDir\\dir2\\file2.txt",
			"test\\test_data\\file2.txt",
		},
	}

	fileList, err = fs.sd.ScanDir(".")
	fs.Require().NoError(err)
	duplicateList = fs.sd.FindDuplicate(fileList)
	fs.Require().Equal(len(duplicateList), 2)
	for key, val := range duplicateList {
		fs.Assert().ElementsMatch(want[key], val)
	}
}
func TestFindTestSuite(t *testing.T) {
	suite.Run(t, new(FindTestSuite))
}

// func TestGetDuplicate(t *testing.T) {
// 	appFS := afero.NewMemMapFs()
// 	appFS.Open(".")
// 	want := [][]string{{
// 		"..\\test\\test_data\\\\testDir\\file2.txt",
// 		"..\\test\\test_data\\\\testDir\\dir2\\file2.txt",
// 		"..\\test\\test_data\\\\testDir\\dir1\\file2.txt",
// 	}}
// 	got, _ := GetDuplicate("..\\test\\test_data\\")
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("GetDuplicate(\".\\test\\test_data\\\" = %q; want %q", got, want)
// 	}
// }

// func ExampleReadDir() {
// 	//fmt.Println(ReadDir(".."))
// 	// Output: [.\\test\\test_data\\testDir\\file2.txt .\\test\\test_data\\testDir\\file2.txt] nil
// }
// [test\test_data\file2.txt test\test_data\testDir\file2.txt test\test_data\testDir\dir1\file2.txt test\test_data\testDir\dir2\file2.txt] appears more times in
// [
// [test\test_data\file2.txt test\test_data\testDir\file2.txt test\test_data\testDir\dir1\file2.txt test\test_data\testDir\dir2\file2.txt]
// [test\test_data\testDir\dir1\file1.txt test\test_data\testDir\dir2\file1.txt test\file1.txt]] than in
// [
// [test\file1.txt test\test_data\testDir\dir2\file1.txt test\test_data\testDir\dir1\file1.txt]
// [test\test_data\testDir\file2.txt test\test_data\testDir\dir1\file2.txt test\test_data\testDir\dir2\file2.txt test\test_data\file2.txt]]
