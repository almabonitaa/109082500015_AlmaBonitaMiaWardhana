# <h1 align="center">Laporan Praktikum Modul 3 - FUNGSI </h1>
<p align="center">[Alma Bonita Mia Wardhana] - [109082500015]</p>

## Unguided 

### 1. [Soal Latihan Modul 3]
#### soal1.go

```go
package main

import (
	"fmt"
)

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_1](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul3/output-soal1.png)
[Program ini menghitung permutasi dan kombinasi dari beberapa angka yang diinput. Fungsi factorial digunakan untuk menghitung faktorial (perkalian dari 1 sampai n). Lalu permutation menghitung banyak susunan (urutan berpengaruh), dan combination menghitung banyak pilihan (urutan tidak berpengaruh). Di main, program membaca 4 angka lalu menampilkan hasil permutasi dan kombinasi untuk dua pasangan angka tersebut.]

### 2. [Soal Latihan Modul 3]
#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	result1 := f(g(h(a)))

	result2 := g(h(f(b)))

	result3 := h(f(g(c)))

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_2](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul3/output-soal2.png)
[Program ini mendefinisikan tiga fungsi sederhana: f(x) = x², g(x) = x - 2, dan h(x) = x + 1. Di dalam main, program membaca tiga input lalu menghitung hasil dari kombinasi fungsi (fungsi di dalam fungsi). result1 menghitung f(g(h(a))), result2 menghitung g(h(f(b))), dan result3 menghitung h(f(g(c))). Terakhir, semua hasil ditampilkan ke layar.]

### 3. [Soal Latihan Modul 3]
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Scan(&x, &y)

	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_3](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/blob/main/modul3/output-soal3.png)
[Program ini digunakan untuk menentukan apakah sebuah titik berada di dalam satu atau dua lingkaran. Fungsi jarak menghitung jarak antara dua titik, sedangkan didalam mengecek apakah jarak titik ke pusat lingkaran kurang dari atau sama dengan radius. Di main, program membaca data dua lingkaran dan satu titik, lalu mengecek posisinya. Hasilnya ditampilkan apakah titik berada di dalam lingkaran 1, lingkaran 2, keduanya, atau di luar keduanya.]