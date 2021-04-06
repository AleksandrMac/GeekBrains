package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"strings"
	"time"
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

		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute)
		defer cancelFunc()
		doneCh := make(chan error)
		messageCh := make(chan string)

		go func(ctx context.Context) {
			// Слушаем сообщения разделённые \n, с таймаутом в 1 минуту
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				doneCh <- err
			}
			messageCh <- message
		}(ctx)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err = <-doneCh:
			return err
		case message := <-messageCh:
			log.Printf("new message received: %s", message)
			// Выполняем обработку сообщения
			responseMessage := strings.ToUpper(message)
			// Возвращаем клиенту обработанное сообщение
			if _, err := conn.Write([]byte(responseMessage + "\n")); err != nil {
				return err
			}
		}
	}
}
