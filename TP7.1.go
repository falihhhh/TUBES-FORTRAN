package main
import "fmt"

func adaDigit(penanda, digit int) bool {
	for penanda > 0 {
		if penanda%10 == digit {
			return true
		}
		penanda /= 10
	}
	return false
}
func solusiValid(angka int) bool {
	if angka <= 0 {
		return false
	}
	banyakDigit := 0
	penanda := 0

	for angka > 0 {
		digit := angka % 10

		if digit == 0 {
			return false
		}
		if adaDigit(penanda, digit) {
			return false
		}
		penanda = penanda*10 + digit
		banyakDigit++
		angka /= 10
	}
	return banyakDigit == 9
}
func tampilkanHasil(hasil bool) {
	fmt.Println(hasil)
}
func main() {
	var namaVariabel, tandaSama string
	var angka int
	fmt.Scan(&namaVariabel, &tandaSama, &angka)
	hasil := solusiValid(angka)
	tampilkanHasil(hasil)
}