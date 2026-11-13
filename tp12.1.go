package main

import "fmt"

const NMAX int = 100

type applicant struct {
	id, name     string
	testGrade    float64
	testDuration int
}
type applicants [NMAX]applicant

func inputData(A *applicants, banyakData *int) {
	var idInput string
	*banyakData = 0
	fmt.Scan(&idInput)
	for idInput != "END" && *banyakData < NMAX {
		A[*banyakData].id = idInput
		fmt.Scan(&A[*banyakData].name, &A[*banyakData].testGrade, &A[*banyakData].testDuration)
		*banyakData++
		fmt.Scan(&idInput)
	}
}
func printData(A applicants, banyakData int) {
	for i := 0; i < banyakData; i++ {
		fmt.Println(A[i].id, A[i].name, A[i].testGrade, A[i].testDuration)
	}
}
func swap(A *applicants, indeks1, indeks2 int) {
	sementara := A[indeks1]
	A[indeks1] = A[indeks2]
	A[indeks2] = sementara
}
func maxPos(A applicants, banyakData, tahap int) int {
	indeksMax := tahap
	for i := tahap + 1; i < banyakData; i++ {
		if A[indeksMax].testGrade < A[i].testGrade {
			indeksMax = i
		} else if A[indeksMax].testGrade == A[i].testGrade {
			if A[indeksMax].testDuration > A[i].testDuration {
				indeksMax = i
			}
		}
	}
	return indeksMax
}
func sortData(A *applicants, banyakData int) {
	for tahap := 0; tahap < banyakData-1; tahap++ {
		indeksMax := maxPos(*A, banyakData, tahap)
		swap(A, indeksMax, tahap)
	}
}
func averageGrade(A applicants, banyakData int) float64 {
	var totalNilai float64 = 0
	for i := 0; i < banyakData; i++ {
		totalNilai += A[i].testGrade
	}
	return totalNilai / float64(banyakData)
}
func averageDuration(A applicants, banyakData int) float64 {
	var totalDurasi int = 0
	for i := 0; i < banyakData; i++ {
		totalDurasi += A[i].testDuration
	}
	return float64(totalDurasi) / float64(banyakData)
}
func main() {
	var daftarPelamar applicants
	var jumlahData int
	inputData(&daftarPelamar, &jumlahData)
	fmt.Println("Data awal:")
	printData(daftarPelamar, jumlahData)
	fmt.Println()
	sortData(&daftarPelamar, jumlahData)
	fmt.Println("Data setelah sortir:")
	printData(daftarPelamar, jumlahData)
	fmt.Println()
	fmt.Printf("Rata-rata nilai tes: %.2f\n", averageGrade(daftarPelamar, jumlahData))
	fmt.Printf("Rata-rata durasi tes: %.2f\n", averageDuration(daftarPelamar, jumlahData))
}
