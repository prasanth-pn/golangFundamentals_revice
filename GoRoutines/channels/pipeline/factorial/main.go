package main

import (
	"fmt"
	"runtime"
)

//change the code above to execute 100 computations concurrently and in  parallel
//use the pipeline pattern  pattern to accomplish this


func main() {
n:=gen()
cf:=factorial(n)
fmt.Println(runtime.NumGoroutine())
for n:=range cf{
	fmt.Println(n)
}
}
func gen()<-chan int{
	out:=make(chan int)
	go func(){
		for i:=0;i<4;i++{
			for j:=3;j<28;j++{
				out<-j
			}
		}
		close(out)
	}()
	return out
}
func factorial(n <-chan int)chan int{
	out :=make(chan int)

	go func ()  {
		for c:=range n{
			out<-fact(c)

		}
		close(out)
	}()
	return out
}

func fact(n int)int{
	total:=1
		for i:=1;i<n;i++{
			total*=i
		}
		return total
}
