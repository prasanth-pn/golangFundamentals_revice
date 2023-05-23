package main

import "fmt"

func main() {
	f:=factorial(14)
	for n:=range f{
		fmt.Println(n)
	}
}
func factorial(n int)chan int64{
	out :=make(chan int64)
	go func ()  {
		var total int64=1
		for i := 1; i <=n; i++ {
			total*=int64(i)
			
		}
		out <-int64(total)
		close(out)
	}()
	return out
}