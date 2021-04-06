package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Print("start")
	sigABRT := make(chan os.Signal, 1)
	signal.Notify(sigABRT, syscall.SIGABRT) //Аварийное завершение
	sigFPE := make(chan os.Signal, 1)
	signal.Notify(sigFPE, syscall.SIGFPE) //Ошибка с плавающей запятой
	sigILL := make(chan os.Signal, 1)
	signal.Notify(sigILL, syscall.SIGILL) //Недопустимая инструкция
	sigINT := make(chan os.Signal, 1)
	signal.Notify(sigINT, syscall.SIGINT) //Сигнал CTRL + C
	sigSEGV := make(chan os.Signal, 1)
	signal.Notify(sigSEGV, syscall.SIGSEGV) //Недопустимый доступ к хранилищу
	sigTERM := make(chan os.Signal, 1)
	signal.Notify(sigTERM, syscall.SIGTERM) //Запрос на завершение

	for {
		select {
		case <-sigABRT:
			gotSIGABRT()
		case <-sigFPE:
			gotSIGFPE()
		case <-sigILL:
			gotSIGILL()
		case <-sigINT:
			gotSIGINT()
		case <-sigSEGV:
			gotSIGSEGV()
		case <-sigTERM:
			gotSIGTERM()
		}
	}
}

func gotSIGABRT() {
	log.Print("SIGABRT handler")
}

func gotSIGFPE() {
	log.Print("SIGFPE handler")
}

func gotSIGILL() {
	log.Print("SIGILL handler")
}

func gotSIGINT() {
	log.Print("SIGINT handler")
}

func gotSIGSEGV() {
	log.Print("SIGSEGV handler")
}

func gotSIGTERM() {
	log.Print("SIGTERM handler")
}
