// package main

// import "fmt"

// func main() {
// 	a := []int{1,1,2,2,2,2}
// 	k := hasGroupsSizeX(a)
// 	fmt.Println(k)
// }

// func hasGroupsSizeX(deck []int) bool {
// 	result:=make(map[int]int)
// 	for _,v:=range deck{
// 		result[v]++
// 	}
// 	res:=result[0]
// 	for _,v:=range result{
// 		res=big(res,v)

// 	}
// 	return res>=1

// }
//
//	func big(x,y int)int{
//		if y==0{
//			return x
//		}
//		return big(x,x%y)
//	}
package main

import "fmt"

func main() {
	a := []int{-3,0,1,-3,1,1,1,-3,10,0}
	
	result:=make(map[int]int)
	for _,v:=range a{
		result[v]++

	}
	for v:=range result{
		fmt.Println("\t",v)
	}
	
}
