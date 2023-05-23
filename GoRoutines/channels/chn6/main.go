package main

import "fmt"
var count int

func main() {
	c:=make(chan int)
	done:=make (chan bool)
	go func(){
		for i := 0; i < 100; i++ {
			c<-i
			count++
		}
		close(c)
	}()
	for i:=0;i<10;i++{

		go func ()  {
			for n:=range c{
				fmt.Println(n)
			}
			done<-true
		}()

	}
	for i:=0;i<10;i++{
		<-done
	}
	fmt.Println(count)
}