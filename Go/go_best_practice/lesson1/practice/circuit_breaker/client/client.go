package main

import (
	"log"
	"net"
	"time"

	"bufio"
	"fmt"
	"os"
)

func main() {
	address := "127.0.0.1:5433"
	timeout := 1 * time.Second
	for {
		// паттерн circuit breaker в простейшей реализации
		// в реальности за функцией writingToServer может скрываться объёмная и сложная логика
		// лишнее выполнение которой мы и хотим прервать
		if isServerAlive(address, timeout) {
			// если сервер стал доступен, возвращаем логику работы с ним
			connectAndWrite(address, timeout)
		}
		// частота опроса проверки жив ли сервер
		time.Sleep(2 * time.Second)
	}
}

func isServerAlive(address string, timeout time.Duration) bool {
	// проверяем можем ли подключиться к серверу
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Printf("can't connect to server, error: %s", err)
		return false
	}
	defer conn.Close()
	return true
}

func connectAndWrite(address string, timeout time.Duration) error {
	// подключаемся к серверу
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Printf("can't connect to server, error: %s", err)
	}
	defer conn.Close()
	log.Printf("connected to server, please write some in console")
	// цикл чтения входных данных от stdin и отправки на сервер
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("can't read sdtin: %s", err)
			return err
		}
		// Отправляем сообщение
		fmt.Fprintf(conn, text+"\n")
		// слушаем ответ
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return err
		}
		log.Printf("Message from server: " + message)
	}
	return nil
}
