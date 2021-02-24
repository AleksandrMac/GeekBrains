package main

import (
	"fmt"
	"os"
	"time"
)

//ErrorWithTime -
type ErrorWithTime struct {
	text string
	time string
}

//New -
func New(text string) error {
	return &ErrorWithTime{
		text: text,
		time: time.Now().Format(time.RFC1123),
	}
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error:\t%s\t%s:", e.text, e.time)
}

func fileCreate(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() error {
		if err := f.Close(); err != nil {
			var err error
			err = New("closed file")
			fmt.Println(fmt.Errorf("%v\n%w", err, err))
			return err
		}
		return nil
	}()
	fmt.Println(f.Name())
	return nil
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			var err error
			err = New("recovered")
			fmt.Printf("%v\n%v\n", err, v)
		}
	}()

	fileCreate("tratata.txt")

	var a int
	a = 1 / a

}
