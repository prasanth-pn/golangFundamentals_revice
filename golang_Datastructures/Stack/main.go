package main

import "fmt"


type Stack struct{
	Items []int
}

func main() {
	stack:=Stack{}
	a:=[]int{23,34,12,45,67,6}
	for _,v:=range a{
		stack.Push(v)
	}
	fmt.Println()
	//stack.Peek()
	stack.Pop()
	fmt.Println(stack.Items)

	
}
func(stack *Stack)Pop(){
	stack.Items=stack.Items[:len(stack.Items)-1]
}
func(stack *Stack)Peek(){
	fmt.Println(stack.Items[len(stack.Items)-1])
}
func(stack *Stack)Push(data int){
	stack.Items=append(stack.Items,data)
}