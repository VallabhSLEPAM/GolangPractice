package main

import "fmt"

//Worker pool that creates 3 workers which takes n numbers and find their respective factorial

type factorial struct {
	num, worker, fact int
}

func main() {

	workers := 3

	inputCh := make(chan factorial)
	outputCh := make(chan factorial)

	for worker := 1; worker <= workers; worker++ {
		go work(inputCh, outputCh, worker)
	}

	go func() {
		for output := range outputCh {
			fmt.Printf("Worker: %d, Num: %d, Factorial: %d\n", output.worker, output.num, output.fact)
		}
	}()

	for num := 1; num <= 10; num++ {
		inputCh <- factorial{
			num: num,
		}
	}
	close(inputCh)

}

func work(inputCh, outputCh chan factorial, worker int) {
	for input := range inputCh {
		input.fact = getFactorial(input.num)
		input.worker = worker
		outputCh <- input
	}
}

func getFactorial(num int) int {

	fact := 1
	for num > 1 {
		fact = fact * num
		num--
	}
	return fact
}
