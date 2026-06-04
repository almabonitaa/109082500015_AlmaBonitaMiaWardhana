# <h1 align="center">Laporan Praktikum Modul 12 - SEARCHING </h1>
<p align="center">[Alma Bonita Mia Wardhana] - [109082500015]</p>

## Unguided 

### 1. [Soal Latihan Modul 12]
#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 12_1](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul12/output-soal1.png)
[Program pertama digunakan untuk menghitung hasil perolehan suara dalam pemilihan ketua RT. Program akan membaca data suara satu per satu hingga menemukan angka 0 sebagai tanda akhir input. Setiap data yang dibaca dihitung sebagai suara masuk, sedangkan hanya angka dari 1 sampai 20 yang dianggap sebagai suara sah karena sesuai dengan nomor calon yang tersedia. Program kemudian menyimpan jumlah suara yang diperoleh setiap calon dan menampilkan total suara masuk, total suara sah, serta daftar calon yang memperoleh suara beserta jumlah suaranya. Dengan cara ini, data yang tidak valid seperti angka negatif atau angka di luar rentang 1–20 dapat diabaikan saat perhitungan suara sah.]

### 2. [Soal Latihan Modul 12]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 12_2](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul12/output-soal2.png)
[Program kedua merupakan pengembangan dari program sebelumnya yang tidak hanya menghitung suara, tetapi juga menentukan siapa yang terpilih menjadi ketua RT dan wakil ketua RT. Setelah seluruh suara dibaca dan divalidasi, program mencari calon dengan jumlah suara terbanyak untuk dijadikan ketua RT. Selanjutnya, program mencari calon dengan jumlah suara terbanyak berikutnya untuk dijadikan wakil ketua. Jika terdapat beberapa calon dengan jumlah suara yang sama, program secara otomatis memilih calon dengan nomor yang lebih kecil karena proses pencarian dilakukan secara berurutan dari nomor terkecil ke terbesar. Di akhir proses, program menampilkan jumlah suara masuk, jumlah suara sah, nomor calon yang menjadi ketua RT, dan nomor calon yang menjadi wakil ketua RT.]

### 3. [Soal Latihan Modul 12]
#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kiri := 0
	kanan := n - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if data[tengah] == k {
			return tengah
		} else if data[tengah] < k {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil := posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 12_3](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul12/output-soal3.png)
[Program ketiga digunakan untuk mencari posisi suatu bilangan dalam kumpulan data yang sudah terurut menaik. Program menerima masukan berupa jumlah data, nilai yang ingin dicari, dan daftar bilangan yang telah tersusun dari kecil ke besar. Data kemudian disimpan ke dalam array melalui prosedur isiArray, sedangkan pencarian dilakukan oleh fungsi posisi menggunakan metode binary search. Metode ini bekerja dengan membandingkan nilai yang dicari dengan elemen di tengah array, sehingga proses pencarian menjadi sangat cepat meskipun jumlah data sangat banyak. Jika bilangan yang dicari ditemukan, program akan menampilkan indeks posisinya yang dihitung mulai dari 0. Namun jika bilangan tersebut tidak ada dalam data, program akan menampilkan tulisan “TIDAK ADA”.]