package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numChan := reader(&wg)
	squareChan := square(numChan, &wg)
	double(squareChan, &wg)
	wg.Wait()
}

func reader(wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			close(out)
			wg.Done()
		}()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() && scanner.Text() != "стоп" {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Printf("Ввод: %d\n", num)
			out <- num
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}()
	return out
}

func square(in chan int, wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			close(out)
			wg.Done()
		}()
		for num := range in {
			val := num * num
			fmt.Println("Квадрат:", val)
			out <- val
		}
	}()
	return out
}

func double(in chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for num := range in {
			val := 2 * num
			fmt.Println("Произведение:", val)
		}
	}()
}
