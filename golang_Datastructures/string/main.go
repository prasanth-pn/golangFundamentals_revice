package main

import "fmt"

func main() {
	result := make(map[int]int)

	result[23] = 123
	result[12] = 124
	for i, v := range result {
		fmt.Println(i, v)
	}
	//print(result)
	for i := 0; i < 3; i++ {
		print(i)
	}
}
