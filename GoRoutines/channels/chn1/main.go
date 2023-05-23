package main

import "fmt"

func main() {
	chn := make(chan int,2)
	for i:=0;i<10;i++{
		go foo(chn, 1+i)
		go foo(chn, 10+i)
	}
		
	for i:=0;i<20;i++{
		fmt.Println(<-chn)
	}

}
func foo(c chan int, value int) {
	c <- value * value
}
