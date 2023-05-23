package main

import "fmt"



type Node struct{
	key string
	value int
	next *Node
}
type HashTable struct{
	table []*Node
	
}
const size int= 100
func main() {
	ht:=HashTable{table: make([]*Node, size)}
	ht.Insert("prasant",456)
	ht.Insert("fazil",45677)
	ht.Insert("ajal",678)
	fmt.Println(ht.Search("fazil"))
	
}

func hash(s string)int{
	var data int
	for _,v:=range s{
		data+=int(v)

	}

	return data%size
}

func (ht *HashTable)Insert(key string,value int){
	index:=hash(key)
	newnode:=&Node{
		key: key,
		value: value,
	}
	if ht.table[index]==nil{
		ht.table[index]=newnode
	}else{
		current:=ht.table[index]
		for current.next!=nil{
			current=current.next
		}
		current.next=newnode
	}
	
}

func(ht *HashTable)Search(key string)int{
	index:=hash(key)

	current:=ht.table[index]

	for current!=nil{
		if current.key==key{
			return current.value

		}
		current=current.next
	}
	return -1
}