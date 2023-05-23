package main

import "fmt"

func main() {
	a := []int{12, 234, 13, 7, 890, 3, 2, 1}
	quicksort(a)
	fmt.Println(a)
}
func quicksort(a []int) {
	if len(a) <= 1 {
		return
	}
	pivot := a[len(a)-1]
	i := 0
	for j := 0; j < len(a); j++ {
		if a[j] < pivot {
			if i != j {
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}
	a[i], a[len(a)-1] = a[len(a)-1], a[i]
	quicksort(a[:i])
	quicksort(a[i+1:])
}
