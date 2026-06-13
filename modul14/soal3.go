package main

import "fmt"

const NMAX = 1000

func insertionSort(A *[NMAX]int, n int) {
	var temp, j int

	for i := 1; i < n; i++ {
		temp = A[i]
		j = i - 1

		for j >= 0 && A[j] > temp {
			A[j+1] = A[j]
			j--
		}

		A[j+1] = temp
	}
}

func cekJarak(A [NMAX]int, n int, jarak *int, tetap *bool) {
	if n <= 1 {
		*tetap = true
		*jarak = 0
		return
	}

	*jarak = A[1] - A[0]
	*tetap = true

	for i := 2; i < n; i++ {
		if A[i]-A[i-1] != *jarak {
			*tetap = false
			return
		}
	}
}

func main() {
	var A [NMAX]int
	var n int
	var x int

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		A[n] = x
		n++
	}

	insertionSort(&A, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(A[i])
	}
	fmt.Println()

	var jarak int
	var tetap bool

	cekJarak(A, n, &jarak, &tetap)

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
