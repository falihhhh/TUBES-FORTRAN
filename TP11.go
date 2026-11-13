package main

import "fmt"

const NMAX = 1001

type TABmurid [NMAX]string

func main() {
	var n, i int
	var m TABmurid
	var murid string
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&m[i])
	}
	fmt.Scan(&murid)
	hasil := CARImurid(m, murid, n)
	if hasil == -1 {
		fmt.Println("Murid tidak terdaftar")
	} else {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d", hasil)
	}
}
func CARImurid(m TABmurid, murid string, n int) int {
	var kiri, kanan, tengah, nemu int
	kiri = 1
	kanan = n - 1
	nemu = -1
	for kiri <= kanan && nemu == -1 {
		tengah = (kiri + kanan) / 2
		if murid < m[tengah] {
			kanan = tengah - 1
		} else if murid > m[tengah] {
			kiri = tengah + 1
		} else {
			nemu = tengah
		}
	}
	return nemu
}
