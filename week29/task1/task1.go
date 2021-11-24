package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func square(num int, wg *sync.WaitGroup) chan int {
	defer wg.Done()
	intChan := make(chan int)
	go func() {
		val := num * num
		fmt.Println("Квадрат:", val)
		intChan <- val
	}()
	return intChan
}

func double(squareChan chan int, wg *sync.WaitGroup) chan int {
	defer wg.Done()
	intChan := make(chan int)
	num, _ := <- squareChan
	go func() {
		val := 2 * num
		fmt.Println("Произведение:", val)
		intChan <- 2 * num
	}()
	return intChan
}

func readLines() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && scanner.Text() != "стоп" {
		var wg sync.WaitGroup
		wg.Add(2)
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Ввод: %d\n", num)
		fc := square(num, &wg)
		sc := double(fc, &wg)
		<- sc
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func main() {
	readLines()
}