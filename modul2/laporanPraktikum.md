# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">[Alma Bonita Mia Wardhana] - [109082500015]</p>

## Unguided 

### 1. [Soal Latihan Modul 2A]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_1](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/modul2/output/output-soal1.png)
[program Go tersebut digunakan untuk menerima tiga input string dari pengguna, lalu menukar urutan nilainya. Program menampilkan urutan awal dari input, kemudian menggunakan variabel temp sebagai penyimpanan sementara untuk menukar posisi data, dan akhirnya menampilkan urutan baru setelah ditukar.]

### 2. [Soal Latihan Modul 2B]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var w1, w2, w3, w4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&w1, &w2, &w3, &w4)

		if w1 != "merah" || w2 != "kuning" || w3 != "hijau" || w4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_2](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/modul2/output/output-soal2.png)
[Program Go tersebut digunakan untuk mengecek apakah urutan warna yang dimasukkan pengguna sudah benar.

Program memberi 5 kali percobaan kepada pengguna untuk memasukkan empat kata. Setiap percobaan, pengguna memasukkan empat warna yang disimpan dalam variabel w1, w2, w3, dan w4. Program kemudian memeriksa apakah urutannya merah, kuning, hijau, dan ungu. Jika ada yang tidak sesuai, maka variabel berhasil akan berubah menjadi false. Setelah lima percobaan selesai, program akan menampilkan hasil akhir apakah percobaan tersebut berhasil (true) atau tidak (false).]

### 3. [Soal Latihan Modul 2C]
#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_3](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/modul2/output/output-soal3.png)
[Program Go tersebut berfungsi untuk menghitung biaya pengiriman parsel berdasarkan beratnya.

Pertama, pengguna diminta memasukkan berat parsel dalam gram. Kemudian program mengubahnya menjadi kilogram (kg) dan sisa gram. Biaya pengiriman dihitung Rp10.000 per kg. Jika beratnya lebih dari 10 kg, maka sisa gram tidak dikenakan biaya. Namun jika 10 kg atau kurang, sisa gram akan dikenakan biaya: Rp5 per gram jika sisa ≥ 500 gram, dan Rp15 per gram jika sisa < 500 gram. Terakhir, program menampilkan detail berat, rincian biaya, dan total biaya pengiriman.]
