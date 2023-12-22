package main

/*first case, go routine is not executed because the main thread finish first*/

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	fmt.Println("main function")
}
