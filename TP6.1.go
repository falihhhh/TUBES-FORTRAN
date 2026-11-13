package main
import "fmt"
func hitungJUMDigit(bilangan int) int {
	if bilangan < 10 {
		return 1
	}
	return 1 + hitungJUMDigit(bilangan/10)
}
func hapusDigitTengahGanjil(bilangan int, posisiSekarang, posisiHapus int) int {
	if bilangan == 0 {
		return 0
	}

	digitTerakhir := bilangan % 10

	if posisiSekarang == posisiHapus {
		return hapusDigitTengahGanjil(bilangan/10, posisiSekarang+1, posisiHapus)
	}

	return hapusDigitTengahGanjil(bilangan/10, posisiSekarang+1, posisiHapus)*10 + digitTerakhir
}
func hapusDigitTengahGenap(bilangan int, posisiSekarang, Phapus1, Phapus2 int) int {
	if bilangan == 0 {
		return 0
	}

	digitTerakhir := bilangan % 10

	if posisiSekarang == Phapus1 || posisiSekarang == Phapus2 {
		return hapusDigitTengahGenap(bilangan/10, posisiSekarang+1, Phapus1, Phapus2)
	}

	return hapusDigitTengahGenap(bilangan/10, posisiSekarang+1, Phapus1, Phapus2)*10 + digitTerakhir
}

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan < 100 {
		fmt.Println("INPUT ERROR")
		return
	}

	panjangBilangan := hitungJUMDigit(bilangan)

	if panjangBilangan%2 == 1 {
		posisiHapus := (panjangBilangan + 1) / 2
		fmt.Println(hapusDigitTengahGanjil(bilangan, 1, posisiHapus))
	} else {
		Phapus1 := panjangBilangan / 2
		Phapus2 := Phapus1 + 1
		fmt.Println(hapusDigitTengahGenap(bilangan, 1, Phapus1, Phapus2))
	}
}