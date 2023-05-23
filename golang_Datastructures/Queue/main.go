package main

import "fmt"
type Queue struct{
	Items []int
}

func main() {
	q:=Queue{}
	a:=[]int{23,45,12,23,67,45}
	for _,v:=range a{
		q.enqueue(v)

	}
	q.dequeue()
	fmt.Println(q.Items)
	
}


//EnQueue >  Adds the variable on last
func(q *Queue)enqueue(data int){
	q.Items=append(q.Items, data)
}
//Dequeue >  takes the first value out
func (q *Queue)dequeue(){
	q.Items=q.Items[1:]
}


