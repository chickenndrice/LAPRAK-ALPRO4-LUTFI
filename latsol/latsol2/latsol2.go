package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	for j := 0; j < n; j++ {
		var r, t float64
		fmt.Scan(&r, &t)

		v := (1.0 / 3.0) * math.Pi * r * r * t
		fmt.Println(v)
	}
}
