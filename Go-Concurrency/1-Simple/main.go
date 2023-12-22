package main

import (
	"fmt"
	"runtime"
)

func sum_up(name string, count_to int, print_every int, done chan bool) {
	my_sum := 0
	for i := 0; i < count_to; i++ {
		if i%print_every == 0 {
			fmt.Printf("%s working on: %d\n", name, i)
		}
		my_sum += 1
	}
	fmt.Printf("%s: %d\n", name, my_sum)
	done <- true
}

func main() {
	runtime.GOMAXPROCS(1)
	done := make(chan bool)

	const COUNT_TO = 10000000
	const PRINT_EVERY = 1000000

	go sum_up("Amy", COUNT_TO, PRINT_EVERY, done)
	go sum_up("Brian", COUNT_TO, PRINT_EVERY, done)

	<-done
	<-done

}
