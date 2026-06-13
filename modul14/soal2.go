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

		data := make([]int, m)

		for i := 0; i < m; i++ {
			fmt.Scan(&data[i])
		}

		selectionSortAsc(data)

		pertama := true

		for i := 0; i < m; i++ {
			if data[i]%2 != 0 {
				if !pertama {
					fmt.Print(" ")
				}
				fmt.Print(data[i])
				pertama = false
			}
		}

		for i := m - 1; i >= 0; i-- {
			if data[i]%2 == 0 {
				if !pertama {
					fmt.Print(" ")
				}
				fmt.Print(data[i])
				pertama = false
			}
		}

		fmt.Println()
	}
}
