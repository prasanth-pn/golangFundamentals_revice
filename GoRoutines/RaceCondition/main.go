package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {

	wg.Add(2)
	go incrementor("foo:")
	go incrementors("ko0:")
	wg.Wait()
	fmt.Println("final counter", counter)

}
func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(1*time.Millisecond))
		counter = x
		fmt.Println(s, i, "counter", counter)
	}
	wg.Done()
}
func incrementors(s string) {
	for i := 0; i < 30; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(500 * time.Microsecond))
		counter = x
		fmt.Println(s, i, "counter", counter)
	}
	wg.Done()
}
