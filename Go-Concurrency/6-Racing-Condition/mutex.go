package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		//mutex must be passed by reference instead by value otherwise each goroutine will have its own mutex instance, and the race condition will remain.
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
