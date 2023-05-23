package main

import "fmt"

func main() {
	a:=[]int{23,1,5,7,34,12,2}
	selectionsort(a)
	fmt.Println(a)
}
func selectionsort(a []int){
	for i:=0;i<len(a);i++{
		min:=i
		for j:=i+1;j<len(a);j++{
			if a[min]>a[j]{
				min=j
			}
		}
		a[min],a[i]=a[i],a[min]
	}
}