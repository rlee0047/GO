package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println("os\t\t", runtime.GOOS)
	fmt.Println("arch\t\t", runtime.GOARCH)
	fmt.Println("cps\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	go f(0)
	fmt.Println("f(0)")
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
