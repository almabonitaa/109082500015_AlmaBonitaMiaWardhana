# <h1 align="center">Laporan Praktikum Modul 5 - REKURSIF </h1>
<p align="center">[Alma Bonita Mia Wardhana] - [109082500015]</p>

## Unguided 

### 1. [Soal Latihan Modul 5]
#### soal1.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_1](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul5/output-soal1.png)
[Program ini digunakan untuk menampilkan deret Fibonacci hingga nilai ke-n yang dimasukkan oleh pengguna, di mana pertama program meminta input angka n, lalu menggunakan perulangan for dari 0 sampai n untuk mencetak setiap nilai Fibonacci, sementara perhitungan dilakukan oleh fungsi fibonacci(n) secara rekursif, yaitu jika n = 0 hasilnya 0, jika n = 1 hasilnya 1, dan jika lebih dari itu maka nilainya didapat dari penjumlahan dua bilangan sebelumnya (fibonacci(n-1) + fibonacci(n-2)), sehingga output yang dihasilkan berupa deret seperti 0 1 1 2 3 5 8 dan seterusnya sesuai input pengguna.]

### 2. [Soal Latihan Modul 5]
#### soal2.go

```go
package main

import "fmt"

func bintang(n int, i int) {
	if i > n {
		return
	}

	for j := 1; j <= i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	bintang(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	bintang(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_2](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul5/output-soal2.png)
[Program ini digunakan untuk menampilkan pola bintang berbentuk segitiga ke bawah sesuai dengan nilai n yang dimasukkan pengguna, di mana program pertama meminta input n, lalu memanggil fungsi bintang(n, 1) yang bekerja secara rekursif untuk mencetak baris demi baris, setiap baris menampilkan jumlah bintang sesuai nilai i (dimulai dari 1 hingga n) menggunakan perulangan for, kemudian setelah mencetak satu baris, fungsi memanggil dirinya sendiri dengan nilai i+1 sampai i lebih besar dari n, sehingga menghasilkan pola bintang bertingkat dari sedikit ke banyak.]

### 3. [Soal Latihan Modul 5]
#### soal3.go

```go
package main

import "fmt"

func faktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	faktor(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_3](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul5/output-soal3.png)
[Program ini digunakan untuk menampilkan semua faktor dari suatu bilangan n yang dimasukkan oleh pengguna, di mana program pertama meminta input n, lalu memanggil fungsi faktor(n, 1) yang bekerja secara rekursif mulai dari i = 1 hingga n, setiap nilai i akan dicek apakah habis membagi n (n % i == 0), jika iya maka i akan ditampilkan sebagai faktor, kemudian fungsi akan terus memanggil dirinya sendiri dengan i+1 sampai i lebih besar dari n, sehingga semua faktor dari n dapat ditampilkan.]

### 4. [Soal Latihan Modul 5]
#### soal4.go

```go
package main

import "fmt"

func pola(n int) {
	if n == 0 {
		return
	}

	fmt.Print(n, " ")

	pola(n - 1)

	if n != 1 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	pola(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_1](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul5/output-soal4.png)
[Program ini digunakan untuk menampilkan pola angka yang turun lalu naik kembali sesuai nilai n yang dimasukkan pengguna, di mana program pertama meminta input n, kemudian memanggil fungsi pola(n) secara rekursif, fungsi akan mencetak nilai n lalu memanggil dirinya sendiri dengan n-1 sampai mencapai 0 sebagai kondisi berhenti, setelah itu saat proses kembali (backtracking), fungsi akan mencetak lagi nilai n kecuali saat n = 1, sehingga menghasilkan pola seperti 5 4 3 2 1 2 3 4 5.]

### 5. [Soal Latihan Modul 5]
#### soal5.go

```go
package main

import "fmt"

func ganjil(i int, n int) {
	if i > n {
		return
	}

	fmt.Print(i, " ")
	ganjil(i+2, n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	ganjil(1, n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_2](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul5/output-soal5.png)
[Program ini digunakan untuk menampilkan deret bilangan ganjil hingga nilai n yang dimasukkan oleh pengguna, di mana program pertama meminta input n, lalu memanggil fungsi `ganjil(1, n)` yang bekerja secara rekursif dimulai dari angka 1, setiap langkah fungsi akan mencetak nilai i kemudian memanggil dirinya sendiri dengan i+2 untuk melompat ke bilangan ganjil berikutnya, dan proses ini akan berhenti ketika i lebih besar dari n, sehingga menghasilkan deret seperti 1 3 5 7 dan seterusnya hingga batas n.]

### 6. [Soal Latihan Modul 5]
#### soal6.go

```go
package main

import "fmt"

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	fmt.Println(pangkat(x, y))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4_3](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul5/output-soal6.png)
[Program ini digunakan untuk menghitung hasil perpangkatan suatu bilangan x dengan pangkat y yang dimasukkan oleh pengguna, di mana program pertama meminta input nilai x dan y, lalu memanggil fungsi pangkat(x, y) yang bekerja secara rekursif, yaitu jika y = 0 maka hasilnya 1 sebagai kondisi dasar, sedangkan jika tidak maka hasilnya adalah x dikalikan dengan hasil pemanggilan fungsi pangkat(x, y-1), sehingga proses ini akan terus berulang sampai nilai y habis dan menghasilkan nilai x pangkat y.]