package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPU's", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	var counter int64

	const g = 100
	var wg sync.WaitGroup
	wg.Add(g)

	for i := 0; i < g; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
		fmt.Println("CPU's", runtime.NumCPU())
	}
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("That's the count", counter)
}
