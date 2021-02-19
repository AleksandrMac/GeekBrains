package main

import (
	"fmt"
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
func main() {
	defer func() {
		if v := recover(); v != nil {
			var err error
			err = New("recovered")
			fmt.Printf("%v\n%v", err, v)
		}
	}()

	var a int
	a = 1 / a
}
