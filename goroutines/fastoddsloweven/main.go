package main

import (
	"fmt"
	"sync"
)

var mySlice []int

func main() {
	num := 16
	signalCh := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go printOdd(signalCh, &wg, num)
	go printEven(signalCh, &wg, num)
	wg.Wait()
}

func printEven(signalCh chan struct{}, wg *sync.WaitGroup, printTill int) {
	for num := 2; num <= printTill; num = num + 2 {

		if _, isOpen := <-signalCh; isOpen {
			println("Even: ", num)
			signalCh <- struct{}{}
		} else {
			println("Even: ", num)
		}

		if num == printTill {
			wg.Done()
		}
	}
}

func printOdd(signalCh chan struct{}, wg *sync.WaitGroup, printTill int) {

	index := 0
	for num := 1; num <= printTill; num = num + 2 {
		if index != 2 {
			index++
			fmt.Println("Odd: ", num)
		} else {
			signalCh <- struct{}{}
			index = 1
			<-signalCh
			fmt.Println("Odd: ", num)

		}
		if num == printTill {
			wg.Done()
		}
	}
	close(signalCh)

}
