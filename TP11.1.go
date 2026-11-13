package main

import "fmt"

const NMAX = 1001

type Penduduk struct {
	nama  string
	tahun int
	kota  string
}
type tabPenduduk [NMAX]Penduduk

func searchTahun(data tabPenduduk, n int, target int) int {
	var left, right, mid, foundIdx int
	left = 0
	right = n - 1
	foundIdx = -1

	for left <= right && foundIdx == -1 {
		mid = (left + right) / 2

		if data[mid].tahun == target {
			foundIdx = mid
		} else if data[mid].tahun < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return foundIdx
}
func main() {
	var n int
	fmt.Scan(&n)
	var data tabPenduduk
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&data[i].nama, &data[i].tahun, &data[i].kota)
	}
	var target int
	fmt.Scan(&target)
	var resultIdx int
	resultIdx = searchTahun(data, n, target)
	if resultIdx != -1 {
		fmt.Printf("%s\n", data[resultIdx].nama)
		fmt.Printf("%d\n", data[resultIdx].tahun)
		fmt.Printf("%s\n", data[resultIdx].kota)
		fmt.Printf("ditemukan di index ke-%d\n", resultIdx)
	} else {
		fmt.Printf("Data Tidak Ditemukan\n")
	}
}
