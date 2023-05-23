package main

import "fmt"

func main() {
	a:=[]int{1,2,3,4,5,6,7,8}
	binarySearch(a,8)
}
func binarySearch(a []int,target int){
	low:=0
	high:=len(a)-1
	for low<=high{
		mid:=low+high/2
		if a[mid]==target{
			fmt.Printf("the data %d is found at the position %d\n",target,mid+1)
			return 
		}else if target >a[mid]{
			low=mid+1
		}else{
			high=mid-1
		}
	}
fmt.Println("data is not found in the array")
	return
}