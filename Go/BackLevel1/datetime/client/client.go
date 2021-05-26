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

	buf := make([]byte, 256) // создаем буфер
	for {
		_, err = conn.Read(buf)
		if err == io.EOF {
			break
		}
		if _, err := io.WriteString(os.Stdout, fmt.Sprintf("Custom output! %s", string(buf))); err != nil { // выводим измененное сообщение сервера в консоль
			log.Fatal(err)
		}
	}
}
