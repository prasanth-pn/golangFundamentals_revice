package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c:=make (chan int)

	go func ()  {
		for i:=0;i<10;i++{
			 fmt.Println("first function",i)
			c<-i
		}
		
		
	}()
	go func ()  {
		for{

			fmt.Println(<-c)
		}
		//fmt.Println("second function")
		
	}()
	go func ()  {
		for{
			fmt.Println(<-c)
		}
		
	}()
	//<-c
	time.Sleep(3*time.Second)

fmt.Println(runtime.NumCPU(),runtime.NumGoroutine())

	
}