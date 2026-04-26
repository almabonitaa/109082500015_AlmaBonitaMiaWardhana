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
