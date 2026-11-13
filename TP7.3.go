package main
import "fmt"

func cekUrutanDigit(angka int, sudahAda7 *bool, sudahAda0 *bool, sudahAda6 *bool) bool {
	if angka == 0 {
		return true
	}
	if !cekUrutanDigit(angka/10, sudahAda7, sudahAda0, sudahAda6) {
		return false
	}
	digit := angka % 10

	if digit == 3 && *sudahAda7 {
		return false
	}
	if digit == 5 && *sudahAda0 {
		return false
	}
	if digit == 2 && *sudahAda6 {
		return false
	}
	if digit == 7 {
		*sudahAda7 = true
	}
	if digit == 0 {
		*sudahAda0 = true
	}
	if digit == 6 {
		*sudahAda6 = true
	}
	return true
}
func validasiBilangan(angka int) {
	var sudahAda7, sudahAda0, sudahAda6 bool

	if cekUrutanDigit(angka, &sudahAda7, &sudahAda0, &sudahAda6) {
		fmt.Println("BILANGAN VALID")
	} else {
		fmt.Println("BILANGAN TIDAK VALID")
	}
}
func main() {
	var namaVariabel, tandaSama string
	var angka int

	fmt.Scan(&namaVariabel, &tandaSama, &angka)
	validasiBilangan(angka)
}