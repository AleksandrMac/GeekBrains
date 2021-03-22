package find_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/AleksandrMac/GeekBrains/Go/Go_level_2/final/find"
)

func TestGetDuplicate(t *testing.T) {
	fmt.Println("123333")
	want := [][]string{{
		"..\\test\\test_data\\\\testDir\\file2.txt",
		"..\\test\\test_data\\\\testDir\\dir2\\file2.txt",
		"..\\test\\test_data\\\\testDir\\dir1\\file2.txt",
	}}
	got, _ := find.GetDuplicate("..\\test\\test_data\\")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetDuplicate(\".\\test\\test_data\\\" = %q; want %q", got, want)
	}
}
