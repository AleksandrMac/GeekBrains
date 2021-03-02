package gorutine

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Task2 Написать программу, которая при получении в канал сигнала SIGTERM
// останавливается не позднее, чем за одну секунду (установить таймаут).
func Task2() {

	ch := make(chan uint16)
	sys := make(chan os.Signal)
	//sys1 := make(chan int)
	errCh := make(chan error)

	signal.Notify(sys, syscall.SIGTERM)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	//go func(ctx context.Context) {
	//time.Sleep(10 * time.Second)
	go func() {
		for mes := range ch {
			fmt.Println(mes)
		}
	}()

	go func(ctx context.Context) {
		for {
			time.Sleep(1 * time.Second)
			if _, ok := ctx.Value("Closed").(string); ok {
				return
			}
			ch <- uint16(time.Now().Nanosecond())
		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			val := <-sys
			fmt.Println(val)
			ctx, cancelFunc = context.WithTimeout(ctx, 2*time.Second)
			ctx = context.WithValue(ctx, "Closed", "true")
			defer cancelFunc()
			var err error
			select {
			case <-ctx.Done():
				close(ch)
				err = ctx.Err()
			case err = <-errCh:
			}
			fmt.Println(err)
		}
	}(ctx)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("call SIGTERM")
		sys <- syscall.SIGTERM
	}()
	//time.Sleep(1 * time.Second)

	//}(ctx)

	var err error
	select {
	// case <-sys:
	// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 	err = ctx.Err()
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-errCh:
	}
	fmt.Println(err)
}
