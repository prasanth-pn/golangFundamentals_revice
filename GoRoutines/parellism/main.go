package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	var sum int
	go foo(sum)
	go koo(sum)
	wg.Wait()

}

func foo(s int) {
	for i := 0; i < 50; i++ {
		fmt.Println("the foo program ", i)
		time.Sleep(0 * time.Second)
		s = s + i
	}
	wg.Done()
}
func koo(s int) {
	for i := 0; i < 50; i++ {
		fmt.Println("the koo", i)
		time.Sleep(0 * time.Second)
		s = s + 1
	}
	wg.Done()
}
