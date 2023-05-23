package main

import "fmt"

func main() {
	a := []int{89,23,233,43,7}
fmt.Println(mergesort(a))
}
func mergesort(a []int)[]int{
	if len(a)<=1{
		return a
	}
	mid:=len(a)/2
	left:=mergesort(a[:mid])
	right:=mergesort(a[mid:])
	return merge(left,right)
}
func merge(left,right []int)[]int{
	result:=[]int{}
	for len(left)>0&&len(right)>0{
		if left[0]<=right[0]{
			result=append(result, left[0])
			left=left[1:]
		}else{
			result=append(result, right[0])
			right=right[1:]
		}
	}
	if len(left)>0{
		result=append(result, left...)
	}
	if len(right)>0{
		result=append(result, right...)
	}
	return result
}