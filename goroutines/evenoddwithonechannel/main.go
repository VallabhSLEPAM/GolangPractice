package main

import (
	"fmt"
	"sync"
)

func main() {

	signalCh := make(chan struct{})
	printTill := 16

	var wg sync.WaitGroup
	wg.Add(1)
	go printEven(signalCh, &wg, printTill)
	go printOdd(signalCh, &wg, printTill)
	wg.Wait()

}

func printEven(signalCh chan struct{}, wg *sync.WaitGroup, num int) {
	signalCh <- struct{}{}
	for even := 2; even <= num; even = even + 2 {
		<-signalCh
		fmt.Println("Even: ", even)
		if num == even {
			wg.Done()
			close(signalCh)
			return
		} else {

			signalCh <- struct{}{}

		}
	}
}

func printOdd(signalCh chan struct{}, wg *sync.WaitGroup, num int) {
	for odd := 1; odd <= num; odd = odd + 2 {
		<-signalCh
		fmt.Println("Odd: ", odd)
		if num == odd {
			wg.Done()
			close(signalCh)
		} else {
			signalCh <- struct{}{}
		}
	}
}
