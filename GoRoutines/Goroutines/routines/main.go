package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var count int

func main() {
	wg.Add(2)

	ch := make(chan int, 50)

	go func(ch <-chan int) {
		for{
		if v ,ok:= <- ch; ok{
			fmt.Println(v)
		}else{
			break
		}
	}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 1232
		ch <- 1233444
		close(ch)
		wg.Done()

	}(ch)
	wg.Wait()
}