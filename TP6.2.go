package main
import "fmt"
func hariKerja(hariMulai, jumlahHari int) int {

	if jumlahHari == 0 {
		return 0
	}
	tambahHariKerja := 0
	if hariMulai >= 1 && hariMulai <= 5 {
		tambahHariKerja = 1
	}
	hariBerikutnya := 1
	if hariMulai == 7 {
		hariBerikutnya = 1
	} else {
		hariBerikutnya = hariMulai + 1
	}
	return tambahHariKerja + hariKerja(hariBerikutnya, jumlahHari-1)
}
func main() {
	var hariMulai, jumlahHari int
	fmt.Scan(&hariMulai, &jumlahHari)

	fmt.Println(hariKerja(hariMulai, jumlahHari))
}

databulat + msbreww