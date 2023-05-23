package main

import "fmt"

func main() {
	a := []int{45, 1, 23, 67, 21, 2}
	insertionsort(a)
	fmt.Println(a)

}

func insertionsort(a []int) { 
	for i := 1; i < len(a); i++ {
		current := i	
					//45, 1, 23, 67, 21, 2
		for j := i - 1; j >= 0; j-- {		//1,45,
			
			if a[j] >= a[current] {
				a[j],a[current]=a[current],a[j]
				current = j
			} else {
				break
			}
		}
	}
}
