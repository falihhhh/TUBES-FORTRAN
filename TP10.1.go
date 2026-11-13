package main

import "fmt"

const NMAX = 999

type Rapot struct {
	matkul string
	nilai  int
}
type TabRapot [NMAX]Rapot

func sequentialSearch(data TabRapot, n int, targetNilai int) {
	statusDitemukan := false

	for i := 0; i < n; i++ {
		if data[i].nilai == targetNilai {
			fmt.Printf("%s %d\n", data[i].matkul, data[i].nilai)
			statusDitemukan = true
		}
	}
	if statusDitemukan {
		fmt.Println("Data ditemukan!")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
func main() {
	var banyakData, nilaiDicari int
	var dataRapot TabRapot

	fmt.Scan(&banyakData)
	for i := 0; i < banyakData; i++ {
		fmt.Scan(&dataRapot[i].matkul, &dataRapot[i].nilai)
	}
	fmt.Scan(&nilaiDicari)
	sequentialSearch(dataRapot, banyakData, nilaiDicari)
}
