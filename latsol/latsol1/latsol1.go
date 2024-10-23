package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Masukkan Angka:")
	fmt.Scan(&n)

	for j := 1; j <= n; j++ {
		sum += j
	}
	fmt.Print(sum)
}
