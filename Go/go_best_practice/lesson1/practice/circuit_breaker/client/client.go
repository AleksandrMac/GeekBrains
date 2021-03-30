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
	//кол-во недачных соединений после которых закрывам канал
	maxBadConnect := 10
	//процент успешных соединений
	succesConnectProcent := 100
	halfOpen := false

	for {
		// паттерн circuit breaker в простейшей реализации
		// в реальности за функцией writingToServer может скрываться объёмная и сложная логика
		// лишнее выполнение которой мы и хотим прервать
		if isServerAlive(address, timeout) {
			maxBadConnect = 10
			log.Println(halfOpen, "   ", maxBadConnect, "     ", succesConnectProcent)
			if halfOpen {

				//для оценки процента успешных соединений рассматриваем последнии maxBadConnect * 2 соединений
				//работаем с целыми числами т.к. этой точности достасточно
				succesConnectProcent += 100 / (maxBadConnect * 2)
				//паузами регулируем плавный возврат к рабочему режиму
				log.Println(succesConnectProcent)
				if succesConnectProcent < 70 {
					time.Sleep(4 * time.Second)
				} else if succesConnectProcent < 90 {
					time.Sleep(3 * time.Second)
				} else {
					//после выхода в рабочей режим снимаем флаг полуоткрытости
					halfOpen = false
					succesConnectProcent = 100
				}
				connectAndWrite(address, timeout)
			} else {
				// если сервер стал доступен, возвращаем логику работы с ним
				connectAndWrite(address, timeout)
			}
		} else {
			log.Println(halfOpen, "   ", maxBadConnect)
			if !halfOpen {
				maxBadConnect--
				if maxBadConnect == 0 {
					halfOpen = true
					succesConnectProcent = 50
				}
			}
		}
		// частота опроса проверки жив ли сервер
		time.Sleep(1 * time.Second)
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
	//return nil
}
