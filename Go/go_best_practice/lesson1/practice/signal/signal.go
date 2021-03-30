package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Print("start")
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT) // регистрируем каналы для получения нотификации указанных сигналов
	sig := <-signalChan // ожидаем получения сигнала из канала
	log.Printf("got %s signal", sig.String())
	log.Print("end")
}
