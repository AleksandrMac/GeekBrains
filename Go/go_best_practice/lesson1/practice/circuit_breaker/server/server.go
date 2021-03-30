package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func main() {
	address := "127.0.0.1:5433"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("server running at address %s", address)

	for {
		if err := waitNewSingleClient(listener); err != nil {
			log.Printf("client connect error: %s", err)
		}
	}
}

func waitNewSingleClient(listener net.Listener) error {
	conn, err := listener.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Print("accept new client connect")
	for {
		// Слушаем сообщения разделённые \n
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return err
		}
		log.Printf("new message received: %s", message)
		// Выполняем обработку сообщения
		responseMessage := strings.ToUpper(message)
		// Возвращаем клиенту обработанное сообщение
		if _, err := conn.Write([]byte(responseMessage + "\n")); err != nil {
			return err
		}
	}
}
