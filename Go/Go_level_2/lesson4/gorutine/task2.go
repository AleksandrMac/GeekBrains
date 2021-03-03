package gorutine

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"
)

// Task2 Написать программу, которая при получении в канал сигнала SIGTERM
// останавливается не позднее, чем за одну секунду (установить таймаут).
func Task2() {
	sys := make(chan os.Signal)
	ch := make(chan int, 10)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(<-ch)
		}
	}()
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		sys <- syscall.SIGTERM
	}()

	select {
	case val := <-sys:
		switch val {
		case syscall.SIGTERM:
			fmt.Println("поступил сигнал SIGTERM :", val.String(), time.Now().Local().String())
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			select {
			case <-ctx.Done():
				fmt.Println("выход по сигналу SIGTERM:", val.String(), time.Now().Local().String())
			}
		default:
			fmt.Println("выход по неизвестному сигналу :", val.String(), time.Now().Local().String())
		}
	}
}
