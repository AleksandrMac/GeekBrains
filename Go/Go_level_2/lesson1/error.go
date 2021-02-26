package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/AleksandrMac/test_version"
)

// ErrorWithTime -
type ErrorWithTime struct {
	text string
	time string
}

// New -
func New(text string) error {
	return &ErrorWithTime{
		text: text,
		time: time.Now().Format(time.RFC1123),
	}
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error:\t%s\t%s:", e.text, e.time)
}

func fileCreate(filename string) (err error) {
	var f *os.File
	f, err = os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = f.Close(); err != nil {
			err = fmt.Errorf("%v\n\t%w", New("closed file"), err)
			fmt.Println("inside", err)
		}
	}()
	fmt.Println(f.Name())
	return
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			var err error
			err = New("recovered")
			fmt.Printf("%v\n%v\n", err, v)
		}
	}()

	fmt.Println("outside", fileCreate("tratata.txt"))

	var a int
	a = 1 / a

}
