package main

import "fmt"

func hitungSkor(waktu [8]int, jumlah *int, total *int) {
	*jumlah = 0
	*total = 0

	for i := 0; i < 8; i++ {
		if waktu[i] < 301 {
			*jumlah++
			*total += waktu[i]
		}
	}
}

func main() {
	var nama string

	var pemenang string
	var maxSoal, minWaktu int
	minWaktu = 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		var waktu [8]int
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var jumlah, total int
		hitungSkor(waktu, &jumlah, &total)

		if jumlah > maxSoal || (jumlah == maxSoal && total < minWaktu) {
			pemenang = nama
			maxSoal = jumlah
			minWaktu = total
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}
