package main

import "fmt"

//pipeline
func main() {
	//setup the pipeline
//c:=gen(2,3,5)
out:=square(gen(2,3,5))
//consume the output
for n:=range out{

	fmt.Println(n)
}
}
func gen(nums ...int)chan int{
	out:=make(chan int)
	go func ()  {
		for _,n:=range nums{
			out<-n
		}
		close(out)
	}()
	return out
}
func square(c chan int)chan int{
	out:=make(chan int)
	go func ()  {
		for n:=range c{
			out <-n*n
		}
		close(out)
	}()
	return out
}