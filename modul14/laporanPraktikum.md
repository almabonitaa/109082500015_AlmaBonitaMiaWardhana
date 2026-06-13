# <h1 align="center">Laporan Praktikum Modul 14 - SORTING </h1>
<p align="center">[Alma Bonita Mia Wardhana] - [109082500015]</p>

## Unguided 

### 1. [Soal Latihan Modul 14]
#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 14_1](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul14/output-soal1.png)
[Program ini digunakan untuk mengurutkan nomor rumah kerabat Hercules pada setiap daerah secara menaik menggunakan algoritma Selection Sort. Program terlebih dahulu membaca jumlah daerah (n), kemudian untuk setiap daerah membaca jumlah rumah (m) dan nomor rumah yang disimpan ke dalam array. Fungsi selectionSortAsc bekerja dengan mencari nilai terkecil pada bagian array yang belum terurut, lalu menukarnya dengan elemen pada posisi saat ini. Proses tersebut dilakukan berulang hingga seluruh data terurut dari yang terkecil ke yang terbesar. Setelah pengurutan selesai, program menampilkan nomor rumah yang sudah terurut pada masing-masing daerah.]

### 2. [Soal Latihan Modul 14]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 14_2](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul14/output-soal2.png)
[Program ini bertujuan menampilkan nomor rumah kerabat dengan urutan yang meminimalkan perpindahan sisi jalan, yaitu menampilkan nomor rumah ganjil terlebih dahulu secara menaik, kemudian nomor rumah genap secara menurun. Data rumah pada setiap daerah dibaca dan disimpan ke dalam array, kemudian seluruh data diurutkan terlebih dahulu menggunakan algoritma Selection Sort secara ascending. Setelah terurut, program mencetak semua bilangan ganjil dari depan ke belakang sehingga urut membesar, lalu mencetak semua bilangan genap dari belakang ke depan sehingga urut mengecil. Hasil akhirnya adalah seluruh nomor ganjil muncul lebih dulu, diikuti nomor genap sesuai ketentuan soal.]

### 3. [Soal Latihan Modul 14]
#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 14_3](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul14/output-soal3.png)
[Program ini digunakan untuk membaca sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non-negatif yang disimpan ke dalam array, sedangkan bilangan negatif berfungsi sebagai tanda berhenti membaca data. Setelah seluruh data masuk, program mengurutkannya menggunakan algoritma Insertion Sort melalui fungsi insertionSort, yaitu dengan mengambil satu elemen dan menyisipkannya ke posisi yang tepat pada bagian array yang sudah terurut. Setelah data terurut, fungsi cekJarak menghitung selisih antara dua elemen pertama sebagai acuan, kemudian memeriksa apakah selisih setiap pasangan elemen berikutnya sama dengan selisih tersebut. Jika semua selisih sama, program menampilkan pesan “Data berjarak x”, sedangkan jika terdapat selisih yang berbeda maka program menampilkan “Data berjarak tidak tetap”.]

### 4. [Soal Latihan Modul 14]
#### soal4.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 14_4](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul14/output-soal4.png)
[Program ini digunakan untuk mengelola data buku dalam sebuah perpustakaan dengan memanfaatkan struktur data Buku yang berisi ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating buku. Fungsi DaftarkanBuku digunakan untuk membaca dan menyimpan seluruh data buku ke dalam array pustaka. Selanjutnya fungsi CetakTerfavorit mencari buku dengan rating tertinggi menggunakan pencarian nilai maksimum dan menampilkan informasi judul, penulis, penerbit, serta tahun terbitnya. Setelah itu fungsi UrutBuku mengurutkan seluruh data berdasarkan rating secara menurun menggunakan algoritma Insertion Sort, sehingga buku dengan rating tertinggi berada di posisi awal array. Fungsi Cetak5Terbaru kemudian menampilkan maksimal lima judul buku dengan rating tertinggi. Terakhir, fungsi CariBuku melakukan pencarian buku berdasarkan rating menggunakan algoritma Binary Search pada array yang telah terurut. Jika rating ditemukan, seluruh informasi buku akan ditampilkan, sedangkan jika tidak ditemukan program akan menampilkan pesan bahwa tidak ada buku dengan rating tersebut.]