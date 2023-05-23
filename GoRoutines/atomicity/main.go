package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)



var wg sync.WaitGroup
var counter int64
func main() {
	wg.Add(2)
	go foo("foo")
	go koo("koo")
	wg.Wait()
	fmt.Println("Final count",counter)
	
}
func foo(s string){
for i := 0; i < 20; i++ {
	atomic.AddInt64(&counter,1)
	fmt.Println(s,i,"counter",counter)
}
wg.Done()
}
func koo(s string){
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&counter,1)
		fmt.Println(s,i,"couter",counter)
		
	}
	wg.Done()  
}