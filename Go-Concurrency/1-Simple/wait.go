package main

/*in this case go routine is executed because of wait method in main thread stops the main execution flow*/

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
