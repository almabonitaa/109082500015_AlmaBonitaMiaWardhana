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
![Screenshot Output Unguided 1_1](https://github.com/almabonitaa/109082500015_AlmaBonitaMiaWardhana/modul2/output/output-soal1.png)
[penjelasan]

