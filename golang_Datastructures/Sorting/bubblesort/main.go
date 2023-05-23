package main

import "fmt"

func main() {
	a:=[]int{23,4523,12,561,78}
	bubblesort(a)
	fmt.Println(a)
	
}
func bubblesort(a []int){
	for i:=0;i<len(a)-1;i++{
		for j:=0;j<len(a)-i-1;j++{
			if a[j]>a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
			}

		}
	}
}