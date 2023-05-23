package main

import (
	"fmt"
	"runtime"
)

func main(){
	c:=increamentor()

	cSum:=puller(c)
	for n:=range cSum{
		fmt.Println(n)
	}
	fmt.Println(runtime.NumCPU(),runtime.NumGoroutine())
}
func increamentor()chan int{
	out:=make(chan int)
	go func ()  {
		for i := 0; i < 10; i++ {
			out<-i
			
		}
		close(out)
		
	}()
	return out
}
func puller(c chan int)chan int{
	out:=make(chan int)
	go func ()  {
		var sum int 
		for n:= range c{
			sum+=n
		}
		out<-sum
		close(out)
	}()
	return out
}