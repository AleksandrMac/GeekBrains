package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		if _, err = io.Copy(os.Stdout, conn); err != io.EOF {
			return
		}
	}()
	if _, err = io.Copy(conn, os.Stdin); err != io.EOF { // until you send ^Z
		return
	}
	fmt.Printf("%s: exit", conn.LocalAddr())
}
