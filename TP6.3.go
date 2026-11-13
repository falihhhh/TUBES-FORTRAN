package main
import "fmt"
func isPeak(bilangan int) bool {

	if bilangan < 100 {
		return false
	}

	digitKanan := bilangan % 10
	digitTengah := (bilangan / 10) % 10
	digitKiri := (bilangan / 100) % 10

	if digitTengah > digitKiri && digitTengah > digitKanan {
		return true
	}
	return isPeak(bilangan / 10)
}

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if isPeak(bilangan) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}