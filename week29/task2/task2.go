package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func squares(ctx context.Context) {
	val := 1
    for {
		select {
			case <- ctx.Done():
				fmt.Println("Выхожу из программы")
				return
			default:
				fmt.Println(val * val)
				val++
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		squares(ctx)
		wg.Done()
	}()
	
	<- termChan
    cancel()
	wg.Wait()
}