package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Consumer struct {
	in chan int
}

func (c Consumer) gracefulShutdown(ctx context.Context) {
	val := 1
	for {
		select {
			case <- ctx.Done():
				fmt.Println("Выхожу из программы")
				close(c.in)
				return
			default:
				c.in <- val
				val++
		}
	}
}

func (c Consumer) square(wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range c.in {
		fmt.Printf("Получено значение: %d; квадрат: %d\n", val, val * val)		
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	workerPoolSize := 100
	consumer := Consumer{
		in: make(chan int, workerPoolSize),
	}

	go consumer.gracefulShutdown(ctx)
	
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	var wg sync.WaitGroup
	wg.Add(workerPoolSize)
	for i := 0; i < workerPoolSize; i++ {
		go consumer.square(&wg)
	}
	
	<- termChan
	cancel()
	wg.Wait()
}