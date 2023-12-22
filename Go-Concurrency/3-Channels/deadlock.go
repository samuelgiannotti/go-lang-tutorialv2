/*
One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.
*/

package main

func main() {
	ch := make(chan int)
	ch <- 5
}
