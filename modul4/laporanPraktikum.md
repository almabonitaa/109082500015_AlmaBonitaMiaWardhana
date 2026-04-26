# <h1 align="center">Laporan Praktikum Modul 4 - PROSEDUR </h1>
<p align="center">[Alma Bonita Mia Wardhana] - [109082500015]</p>

## Unguided 

### 1. [Soal Latihan Modul 4]
#### soal1.go

```go
package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutasi(a, c), kombinasi(a, c))

	fmt.Println(permutasi(b, d), kombinasi(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_1](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul4/output-soal1.png)
[Program ini dibuat untuk menghitung dua hal dari angka yang dimasukkan pengguna. Pertama, program membaca empat angka yaitu a, b, c, dan d. Lalu program punya fungsi untuk menghitung faktorial dengan cara mengalikan angka dari 1 sampai angka tersebut. Setelah itu, ada dua fungsi lain yang menggunakan hasil faktorial tadi untuk mencari hasil dari pasangan a dengan c, dan b dengan d. Hasilnya kemudian ditampilkan dalam dua baris, baris pertama untuk a dan c, dan baris kedua untuk b dan d. Secara sederhana, program ini hanya mengambil input, mengolahnya dengan beberapa fungsi, lalu menampilkan hasilnya.]

### 2. [Soal Latihan Modul 4]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_2](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul4/output-soal2.png)
[Program ini digunakan untuk menentukan pemenang dari beberapa peserta berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu yang dibutuhkan. Pertama, program membaca nama peserta dan 8 waktu pengerjaan soal, lalu menghitung berapa soal yang selesai (hanya yang waktunya kurang dari 301) dan menjumlahkan total waktunya melalui fungsi hitungSkor. Program akan terus membaca data sampai input “Selesai”. Selama proses itu, program membandingkan setiap peserta untuk mencari yang terbaik, yaitu yang menyelesaikan soal paling banyak, dan jika sama, yang total waktunya paling kecil. Di akhir, program menampilkan nama pemenang beserta jumlah soal yang diselesaikan dan total waktunya.]

### 3. [Soal Latihan Modul 4]
#### soal3.go

```go
package main

import "fmt"

func cetakDeret(n int) {
	for {
		fmt.Print(n, " ")

		if n == 1 {
			break
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var n int

	fmt.Scan(&n)

	cetakDeret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_3](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul4/output-soal3.png)
[Program ini digunakan untuk menampilkan deret angka berdasarkan suatu nilai awal yang dimasukkan pengguna. Setelah membaca angka n, program memanggil fungsi cetakDeret yang akan mencetak angka tersebut lalu terus mengubah nilainya mengikuti aturan: jika genap dibagi 2, dan jika ganjil dikali 3 lalu ditambah 1. Proses ini dilakukan berulang dalam perulangan sampai nilai n menjadi 1, dan setiap perubahan nilai langsung ditampilkan dalam satu baris. Jadi, program ini sederhana, hanya menerima input lalu menampilkan urutan angka sampai berhenti di 1.]