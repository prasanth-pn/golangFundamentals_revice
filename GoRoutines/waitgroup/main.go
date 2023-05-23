package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("program started")
}
func main() {
	//wg.Add(2)
	koo()

	 foo()
	//wg.Wait()

}
func foo(){
	for i:=0;i<10;i++{

		fmt.Println("hi foo program  ")
		time.Sleep(2*time.Second)
	}
	//wg.Done()
}
func koo(){
	for i:=0;i<10;i++{

		fmt.Println("hi koo")
		time.Sleep(1*time.Second)
		
	}
	//wg.Done()
}
