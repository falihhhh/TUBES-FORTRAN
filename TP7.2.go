package main
import "fmt"

func cariRepunit(k, panjang, sisaLama int) int {
	if panjang > k {
		return -1
	}
	sisaBaru := (sisaLama*10 + 1) % k

	if sisaBaru == 0 {
		return panjang
	}
	return cariRepunit(k, panjang+1, sisaBaru)
}
func panjangRepunitMinimum(k int) int {
	if k <= 0 {
		return -1
	}
	return cariRepunit(k, 1, 0)
}
func tampilkanHasil(hasil int) {
	fmt.Println(hasil)
}
func main() {
	var namaVariabel, tandaSama string
	var nilaiK int
	fmt.Scan(&namaVariabel, &tandaSama, &nilaiK)
	hasil := panjangRepunitMinimum(nilaiK)
	tampilkanHasil(hasil)
}