package main
import "fmt"

const NMAX int = 100
type peserta struct {
	id     string
	nama   string
	nilai  int
	durasi int
}
type tabPeserta [NMAX]peserta
func isiArray(T *tabPeserta, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&T[i].id, &T[i].nama, &T[i].nilai, &T[i].durasi)
	}
}
func tampilPeserta(p peserta) {
	fmt.Println(p.id, p.nama, p.nilai, p.durasi)
}
func insertionSort(T *tabPeserta, n int) {
	var pass, i int
	var temp peserta
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = T[pass]
		for i > 0 &&
			(temp.nilai > T[i-1].nilai ||
				(temp.nilai == T[i-1].nilai &&
					temp.durasi < T[i-1].durasi)) {
			T[i] = T[i-1]
			i--
		}
		T[i] = temp
		pass++
	}
}
func rataNilai(T tabPeserta, n int) float64 {
	var i, total int
	for i = 0; i < n; i++ {
		total += T[i].nilai
	}
	return float64(total) / float64(n)
}
func hitungAtasRata(T tabPeserta, n int, rata float64) int {
	var i, jml int
	for i = 0; i < n; i++ {
		if float64(T[i].nilai) > rata {
			jml++
		}
	}
	return jml
}
func main() {
	var T tabPeserta
	var n, i, jml int
	var rata float64
	fmt.Scan(&n)
	isiArray(&T, n)
	insertionSort(&T, n)
	fmt.Println("Data setelah diurutkan:")
	for i = 0; i < n; i++ {
		tampilPeserta(T[i])
	}
	fmt.Println()
	fmt.Println("Peserta terbaik:")
	tampilPeserta(T[0])
	rata = rataNilai(T, n)
	jml = hitungAtasRata(T, n, rata)
	fmt.Println()
	fmt.Println("Jumlah peserta di atas rata-rata:", jml)
}