package main

import "fmt"

const NMAX = 7919

type Buku struct {
	ID        string
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku [NMAX]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].ID,
			&pustaka[i].Judul,
			&pustaka[i].Penulis,
			&pustaka[i].Penerbit,
			&pustaka[i].Eksemplar,
			&pustaka[i].Tahun,
			&pustaka[i].Rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}

	idxMax := 0

	for i := 1; i < n; i++ {
		if pustaka[i].Rating > pustaka[idxMax].Rating {
			idxMax = i
		}
	}

	fmt.Println("Buku Terfavorit")
	fmt.Println("Judul    :", pustaka[idxMax].Judul)
	fmt.Println("Penulis  :", pustaka[idxMax].Penulis)
	fmt.Println("Penerbit :", pustaka[idxMax].Penerbit)
	fmt.Println("Tahun    :", pustaka[idxMax].Tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var j int

	for i := 1; i < n; i++ {
		temp = pustaka[i]
		j = i - 1

		for j >= 0 && pustaka[j].Rating < temp.Rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := 5

	if n < 5 {
		batas = n
	}

	fmt.Println("5 Buku Rating Tertinggi")

	for i := 0; i < batas; i++ {
		fmt.Println(pustaka[i].Judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 0
	kanan := n - 1
	ketemu := -1

	for kiri <= kanan && ketemu == -1 {
		tengah := (kiri + kanan) / 2

		if pustaka[tengah].Rating == r {
			ketemu = tengah
		} else if r > pustaka[tengah].Rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ketemu == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		fmt.Println("Data Buku Ditemukan")
		fmt.Println("Judul     :", pustaka[ketemu].Judul)
		fmt.Println("Penulis   :", pustaka[ketemu].Penulis)
		fmt.Println("Penerbit  :", pustaka[ketemu].Penerbit)
		fmt.Println("Tahun     :", pustaka[ketemu].Tahun)
		fmt.Println("Eksemplar :", pustaka[ketemu].Eksemplar)
		fmt.Println("Rating    :", pustaka[ketemu].Rating)
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}
