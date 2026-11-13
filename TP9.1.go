package main 
import "fmt"
type game struct {
    nama string
    populasi float64
    nilai int
}
type TABgame [20]game

func main(){
    var n, k int 
    var data TABgame
    fmt.Scan(&n, &k)

    for i := 0; i < n; i++ {
        fmt.Scan(&data[i].nama)
        data[i].populasi = float64(k) * 0.10
        pendudukKOTA := int(data[i].populasi)
        k = k - pendudukKOTA
        data[i].nilai = pendudukKOTA % 1000
    }
    for i := 0; i < n; i++ {
        LEVELkejahatan := float64(data[i].nilai) * 2.5 / 100.0

        if LEVELkejahatan > 5 {
            fmt.Printf("%s dengan tingkat kejahatan: %.2f%%\n", data[i].nama, LEVELkejahatan)
        }
    }
}