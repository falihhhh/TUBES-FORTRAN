package main

import "fmt"

const NMAX int = 9999

type arrString [NMAX]string

func bacaData(A *arrString, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}
func cetakData(A arrString, n int) {
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Printf("%s.\n", A[i])
		} else {
			fmt.Printf("%s, ", A[i])
		}
	}
}
func swap(A *arrString, indeks1, indeks2 int) {
	sementara := A[indeks1]
	A[indeks1] = A[indeks2]
	A[indeks2] = sementara
}
func minPos(A arrString, n, tahap int) int {
	indeksMin := tahap
	for i := tahap + 1; i < n; i++ {
		if A[i] < A[indeksMin] {
			indeksMin = i
		}
	}
	return indeksMin
}
func maxPos(A arrString, n, tahap int) int {
	indeksMax := tahap
	for i := tahap + 1; i < n; i++ {
		if A[i] > A[indeksMax] {
			indeksMax = i
		}
	}
	return indeksMax
}
func selectionSortAscend(A *arrString, n int) {
	for tahap := 0; tahap < n-1; tahap++ {
		indeksMin := minPos(*A, n, tahap)
		swap(A, indeksMin, tahap)
	}
}
func selectionSortDescend(A *arrString, n int) {
	for tahap := 0; tahap < n-1; tahap++ {
		indeksMax := maxPos(*A, n, tahap)
		swap(A, indeksMax, tahap)
	}
}
func main() {
	var daftarNama arrString
	var jumlahData int
	fmt.Scan(&jumlahData)
	bacaData(&daftarNama, jumlahData)
	selectionSortAscend(&daftarNama, jumlahData)
	fmt.Println("Data setelah diurutkan secara Ascending:")
	cetakData(daftarNama, jumlahData)
	selectionSortDescend(&daftarNama, jumlahData)
	fmt.Println("Data setelah diurutkan secara Descending:")
	cetakData(daftarNama, jumlahData)
}
