package main

import "fmt"

func main() {
	c:=incrementor()
	sum:=puller(c)
	for n:=range sum{
		fmt.Println(n)
	}
}
func incrementor()<-chan int{
	out :=make(chan int)
	go func ()  {
		for i := 0; i < 10; i++ {
			out<-i
			
		}
		close(out)
	}()
	return out
}
func puller(c <-chan int)<-chan int{
	out := make(chan int)
	sum:=0
go func ()  {
	for v:=range c{
		sum+=v
	}
	out<-sum
	close(out)
}()
return out

}
/*
The optional <-operator specific the channel direction send or recieve  if no direction is given , the channel is bidirectional.
*/