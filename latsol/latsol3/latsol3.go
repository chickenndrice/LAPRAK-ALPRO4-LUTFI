package main

import "fmt"

func main() {
	var a, b, result int
	fmt.Scan(&a, &b)

	result = 1
	for j := 0; j < b; j++ {
		result *= a
	}

	fmt.Println(result)
}
