package main

import "fmt"

const NMAX = 100

type student struct {
	nim, nama string
	nilai     float64
}
type daftarSTUDENT [NMAX]student

func CEKduplikat(data daftarSTUDENT, n int, NIMcari string) bool {
	for i := 0; i < n; i++ {
		if data[i].nim == NIMcari {
			return true
		}
	}
	return false
}
func CLEARdata(data daftarSTUDENT, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %g\n", data[i].nim, data[i].nama, data[i].nilai)
	}
}
func main() {
	var dataSTUDENT daftarSTUDENT
	var BYKdata int = 0
	var inputNIM, inputNAMA string
	var inputNILAI float64

	for {
		fmt.Scan(&inputNIM)
		if inputNIM == "STOP" {
			break
		}
		fmt.Scan(&inputNAMA, &inputNILAI)
		if CEKduplikat(dataSTUDENT, BYKdata, inputNIM) == false {
			dataSTUDENT[BYKdata].nim = inputNIM
			dataSTUDENT[BYKdata].nama = inputNAMA
			dataSTUDENT[BYKdata].nilai = inputNILAI
			BYKdata++
		}
	}
	CLEARdata(dataSTUDENT, BYKdata)
}
