package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	go func() {
		defer c.Close()
		for {
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		// копирование сообщений в открытый канала
		for {
			if _, err := io.Copy(c, os.Stdin); err != nil {
				return
			}
		}
	}()
}
