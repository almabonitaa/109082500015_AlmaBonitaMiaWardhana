package main

import "fmt"

func ganjil(i int, n int) {
	if i > n {
		return
	}

	fmt.Print(i, " ")
	ganjil(i+2, n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	ganjil(1, n)
}
