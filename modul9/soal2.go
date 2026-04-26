package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
	fmt.Scan(&idx)

	arr = append(arr[:idx], arr[idx+1:]...)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	sum := 0.0
	for i := 0; i < len(arr); i++ {
		sum += float64(arr[i])
	}
	mean := sum / float64(len(arr))
	fmt.Println(mean)

	var variance float64
	for i := 0; i < len(arr); i++ {
		diff := float64(arr[i]) - mean
		variance += diff * diff
	}
	variance /= float64(len(arr))
	stddev := math.Sqrt(variance)
	fmt.Println(stddev)

	var target int
	fmt.Scan(&target)

	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			count++
		}
	}
	fmt.Println(count)
}
