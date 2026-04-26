# <h1 align="center">Laporan Praktikum Modul 4 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">[Alma Bonita Mia Wardhana] - [109082500015]</p>

## Unguided 

### 1. [Soal Latihan Modul 9]
#### soal1.go

```go
package main

import "fmt"

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func dalamLingkaran(t Titik, l Lingkaran) bool {
	dx := t.x - l.pusat.x
	dy := t.y - l.pusat.y
	return dx*dx+dy*dy <= l.r*l.r
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	l1 := Lingkaran{Titik{cx1, cy1}, r1}
	l2 := Lingkaran{Titik{cx2, cy2}, r2}
	t := Titik{x, y}

	in1 := dalamLingkaran(t, l1)
	in2 := dalamLingkaran(t, l2)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 9_1](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul9/output-soal1.png)
[Kode tersebut adalah program sederhana dalam bahasa Go yang digunakan untuk menentukan posisi sebuah titik terhadap dua lingkaran. Pertama, dibuat dua tipe data yaitu Titik untuk menyimpan koordinat (x, y) dan Lingkaran untuk menyimpan titik pusat serta jari-jari. Fungsi dalamLingkaran digunakan untuk mengecek apakah suatu titik berada di dalam lingkaran dengan menghitung jarak titik ke pusat lingkaran (tanpa akar, cukup dibandingkan kuadratnya dengan jari-jari). Di dalam main, program membaca input untuk dua lingkaran dan satu titik, lalu memeriksa apakah titik tersebut berada di dalam lingkaran pertama, kedua, keduanya, atau di luar keduanya. Hasilnya kemudian ditampilkan dalam bentuk kalimat yang sesuai, sehingga pengguna bisa langsung tahu posisi titik tersebut.]

### 2. [Soal Latihan Modul 9]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 9_2](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul9/output-soal2.png)
[Kode ini adalah program Go yang bekerja dengan array untuk mengolah sekumpulan angka. Pertama, program membaca jumlah data lalu mengisi array dengan angka-angka dari input. Setelah itu, program menampilkan seluruh isi array, lalu menampilkan elemen dengan indeks ganjil dan indeks genap secara terpisah. Selanjutnya, program meminta angka x untuk menampilkan elemen pada indeks kelipatan x, kemudian meminta indeks tertentu untuk dihapus dari array dan menampilkan hasil array yang sudah diperbarui. Setelah itu, program menghitung rata-rata dari data yang tersisa, lalu menghitung standar deviasi untuk melihat seberapa jauh penyebaran datanya. Terakhir, program meminta sebuah angka target dan menghitung berapa kali angka tersebut muncul di dalam array.]

### 3. [Soal Latihan Modul 9]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Scan(&klubA)
	fmt.Scan(&klubB)

	var skorA, skorB int
	hasil := []string{}
	i := 1

	for {
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}

		i++
	}

	for j := 0; j < len(hasil); j++ {
		fmt.Printf("Hasil %d : %s\n", j+1, hasil[j])
	}

	fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 9_3](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul9/output-soal3.png)
[Kode ini adalah program Go yang digunakan untuk mencatat hasil pertandingan antara dua klub. Pertama, program meminta nama dua klub, lalu secara berulang menerima input skor dari masing-masing klub di setiap pertandingan. Proses input akan berhenti jika salah satu skor bernilai negatif (sebagai tanda selesai). Untuk setiap pertandingan, program akan menentukan pemenangnya: jika skor klub A lebih besar maka klub A menang, jika klub B lebih besar maka klub B menang, dan jika sama maka hasilnya “Draw”. Semua hasil tersebut disimpan dalam sebuah array. Setelah input selesai, program menampilkan hasil setiap pertandingan secara berurutan, lalu diakhiri dengan pesan bahwa pertandingan telah selesai.]

### 4. [Soal Latihan Modul 9]
#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		if ch != ' ' && ch != '\n' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Teks : ")
	cetakArray(tab, m)

	temp := tab
	balikanArray(&temp, m)

	fmt.Print("Reverse teks : ")
	cetakArray(temp, m)

	fmt.Print("Palindrom ? ")
	fmt.Println(palindrome(tab, m))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 9_4](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul9/output-soal4.png)
[Kode ini adalah program Go yang digunakan untuk mengolah teks menjadi array karakter, lalu mengecek apakah teks tersebut termasuk palindrom. Program mendefinisikan tipe tabel sebagai array karakter (rune) dengan kapasitas maksimal 127. Fungsi isiArray membaca input karakter satu per satu sampai menemukan tanda titik (.) sebagai akhir input, sambil mengabaikan spasi dan enter. Kemudian fungsi cetakArray menampilkan isi array, sedangkan balikanArray digunakan untuk membalik urutan karakter dalam array. Fungsi palindrome berfungsi mengecek apakah teks sama jika dibaca dari depan dan belakang. Di dalam main, program membaca teks, menampilkannya, menampilkan versi terbalik, lalu memberikan hasil apakah teks tersebut termasuk palindrom atau tidak, sehingga pengguna bisa memahami struktur dan sifat teks yang dimasukkan.]