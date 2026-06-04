package main

import "fmt"

func main() {
	var suara [21]int
	var pilih int
	var totalMasuk, totalSah int

	for {
		fmt.Scan(&pilih)

		if pilih == 0 {
			totalMasuk++
			break
		}

		totalMasuk++

		if pilih >= 1 && pilih <= 20 {
			suara[pilih]++
			totalSah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", totalSah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}
