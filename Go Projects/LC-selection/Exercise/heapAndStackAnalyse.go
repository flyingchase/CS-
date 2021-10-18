package main

import (
	"fmt"
	"runtime/debug"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Example(slice []string, str string, i int) {
	debug.PrintStack()
}
