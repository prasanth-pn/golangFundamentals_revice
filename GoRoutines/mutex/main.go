package main

import (
	"fmt"
	"sync"
)


var mutex sync.Mutex
var wg sync.WaitGroup
var count int
func main() {
	wg.Add(2)
	go foo("foo")
	go koo("koo")
	wg.Wait()
fmt.Println("final couter",count)	
}
func foo(s string){
	for i := 0; i < 20; i++ {
		mutex.Lock()
		x:=count
		x++
		count=x
		fmt.Println(s,i,"couter",count)
		mutex.Unlock()
		
	}
	wg.Done()

}
func koo(s string){
	for i := 0; i < 20; i++ {
		mutex.Lock()
		x:=count
		x++
		count=x   
		fmt.Println(s,i,"count",count)
		mutex.Unlock()
		
	}
	wg.Done()
}