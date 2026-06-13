package main

import "fmt"

func selectionSortAsc(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for d := 0; d < n; d++ {
		var m int
		fmt.Scan(&m)

		rumah := make([]int, m)

		for i := 0; i < m; i++ {
			fmt.Scan(&rumah[i])
		}

		selectionSortAsc(rumah)

		for i := 0; i < m; i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[i])
		}
		fmt.Println()
	}
}
