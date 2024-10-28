package main

import (
	"context"
	"fmt"
	"time"
)

/*
Program with 2 goroutine, one which prints after every 300ms and other which prints after every 500ms. All of this should happen for 5 Seconds
*/
func main() {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goRoutine(ctx, 300, ch1, "First Goroutine")
	go goRoutine(ctx, 500, ch2, "Second Goroutine")

	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting ...")
			return
		case returnedValue, isOpen := <-ch1:
			if !isOpen {
				ch1 = nil
			} else {
				fmt.Println(returnedValue)
			}
		case returnedValue, isOpen := <-ch2:
			if !isOpen {
				ch2 = nil
			} else {
				fmt.Println(returnedValue)
			}
		}
	}

}

func goRoutine(ctx context.Context, timeInMS int, ch1 chan string, message string) {
	ticker := time.NewTicker(time.Duration(timeInMS * int(time.Millisecond)))
	defer ticker.Stop()
	defer close(ch1)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ch1 <- message
		}
	}
}
