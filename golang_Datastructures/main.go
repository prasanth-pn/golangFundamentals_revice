package main

import (
	"sync"
	"time"
)
var wg sync.WaitGroup
func main() {
	wg.Add(1)
	// go func (){
	// 	time.Sleep(time.Second*2)
	// 	//wg.Wait()
	// }()
	go foo()
	wg.Wait()
	
}
func foo(){
	time.Sleep(time.Second*3)

}
