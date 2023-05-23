package main

import "fmt"

type LinkedList struct {
	Head   *Node
	Length int
}
type Node struct {
	Data int
	Next *Node
}

func main() {
	list := LinkedList{}
	a := []int{34, 56, 23, 12, 67, 89, 54, 32222}
	for _, n := range a {
		list.Create(n)
	}
	list.Traverse()
	fmt.Println("------------------------------------------------")
	// list.Reverse()
	 list.Find(32222)
	 //list.Delete(12)
	 //list.Delete(32222)
	list.DeleteEnd()
	list.Traverse()

}
func (list *LinkedList)Find(data int){
	current:=list.Head
	if current.Data==data{
		fmt.Println("The data in the head position")
		return
	}
	for current.Next!=nil{

		if current.Next.Data==data{
			fmt.Println("The data is present in the list")
			return 
		} 
		current=current.Next
	}
	fmt.Println("The data is not present")

}
func (list *LinkedList) DeleteEnd() {

	if list.Head==nil{
		return
	}
	if list.Head.Next==nil{
		list.Head=nil
		return 
	}
	prev:=list.Head

	current:=list.Head.Next

	for current.Next!=nil{
		prev=current
		current=current.Next
	}
	prev.Next=nil

}
func (list *LinkedList) Delete(data int) {
	if list.Head == nil {
		fmt.Println("list is empty")
		return
	}
	if list.Head.Data == data {
		list.Head = list.Head.Next
	} else {
		current := list.Head.Next
		for current.Next != nil {
prev:=current
			if current.Next.Data == data {
				prev.Next = current.Next.Next
			}
			current = current.Next
		}
	}
}
func (list *LinkedList) Create(data int) {
	newnode := new(Node)
	newnode.Data = data
	newnode.Next = nil
	if list.Head == nil {
		list.Head = newnode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newnode
	}
}
func (list *LinkedList) Traverse() {
	if list.Head == nil {
		fmt.Println("no data available")
		return
	}
	current := list.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
func (list *LinkedList) Reverse() {
	if list.Head == nil {
		fmt.Println("list is empty")
		return 
	} else {
		current := list.Head
		var prev, next *Node
		for current != nil {
			next = current.Next
			current.Next = prev
			prev = current
			current = next
		}
		list.Head = prev
	}

}
