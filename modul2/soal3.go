package main

import "fmt"

func main() {
	var beratGram int
	var kg, sisaGram int
	var biayaKg, biayaSisa, totalBiaya int

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&beratGram)

	kg = beratGram / 1000
	sisaGram = beratGram % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}

	totalBiaya = biayaKg + biayaSisa

	fmt.Println("Detail berat:", kg, "kg +", sisaGram, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}
