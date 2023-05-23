package main

import "fmt"
var count int

func main() {
//	n := 10

	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				count++
				c <- i

			}

			done <- true
		}()
	}
		go func() {
			for i := 0; i < 10; i++ {
				<-done
			}
			close(c)
		}()
		for v := range c {
			fmt.Println(v)
		}
		fmt.Println(count)
	
}
