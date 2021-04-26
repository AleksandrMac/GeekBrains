package gorutine

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//nolint:gosimple
// Task2 Написать программу, которая при получении в канал сигнала SIGTERM
// останавливается не позднее, чем за одну секунду (установить таймаут).
func Task2() {
	var (
		sys         = make(chan os.Signal)
		ctx, cancel = context.WithCancel(context.Background())
		hr          = func(ctx context.Context, cancel context.CancelFunc) {
			select {
			case <-ctx.Done():
				close(sys)
				cancel()
				fmt.Println("goes home")
				return
			case val := <-sys:
				switch val {
				case syscall.SIGTERM:
					cancel()
					fmt.Println("поступил сигнал SIGTERM :", val.String(), time.Now().Local().String())
					time.Sleep(1 * time.Second)
					os.Exit(int(syscall.SIGTERM))
				default:
					cancel()
					fmt.Println("поступил неизвестный сигнал выход по неизвестному сигналу :", val.String(), time.Now().Local().String())
					time.Sleep(1 * time.Second)
					os.Exit(int(syscall.SIGKILL))
				}
			}
		}
		jobs    = make(chan int)
		manager = func(ctx context.Context) {
			for job := 0; ; job++ {
				select {
				case <-ctx.Done():
					close(jobs)
					fmt.Println("manager goes home")
					return
				default:
					fmt.Printf("manager create job %d\n", job)
					jobs <- job
				}
			}
		}

		chanSize uint32 = 10
		resource        = make(chan struct{}, chanSize)
		worker          = func(id int) {
			defer func() { <-resource }()
			for job := range jobs {
				fmt.Printf("worker %d starts processing of %d\n", id, job)
				<-time.NewTicker(1 * time.Second).C
				fmt.Printf("worker %d completes processing of %d\n", id, job)
			}
			fmt.Printf("worker %d goes home\n", id)
		}
	)
	signal.Notify(sys, syscall.SIGINT, syscall.SIGTERM)
	go manager(ctx)
	go hr(ctx, cancel)

	for i := 0; i < cap(resource); i++ {
		resource <- struct{}{}
		go worker(i)
	}

	select {
	case <-ctx.Done():
		for i := 0; i < cap(resource); i++ {
			resource <- struct{}{}
		}
		close(resource)
		return
	}
}
