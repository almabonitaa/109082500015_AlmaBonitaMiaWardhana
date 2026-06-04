package main

import "fmt"

func main() {
	var suara [21]int
	var pilih int
	var totalMasuk, totalSah int

	for {
		fmt.Scan(&pilih)

		totalMasuk++

		if pilih == 0 {
			break
		}

		if pilih >= 1 && pilih <= 20 {
			suara[pilih]++
			totalSah++
		}
	}

	ketua := 1
	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	wakil := -1
	for i := 1; i <= 20; i++ {
		if i != ketua {
			wakil = i
			break
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			if suara[i] > suara[wakil] {
				wakil = i
			}
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", totalSah)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
