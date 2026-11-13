package main 
import "fmt"
type item struct {
    nama string
    harga int
}
type TABitem [100]item

func main () {
    var data TABitem
    var n, IDXmurah, IDXmahal int
    
    fmt.Scan(&n)
     for i := 0; i < n; i++ {
        fmt.Scan(&data[i].nama)
        fmt.Scan(&data[i].harga)
     }
     IDXmurah = 0
     IDXmahal = 0

     for i := 0; i < n; i++ {
        if data[i].harga > data[IDXmahal].harga {
            IDXmahal = i
        }
         if data[i].harga > data[IDXmurah].harga {
            IDXmurah = i
        }
    }
    fmt.Println("Nama Item\t\t Harga")

    for i := 0; i < n; i++ {
        fmt.Printf("%s\t\t%d\n", data[i].nama, data[i].harga)
    }
    fmt.Println()
    fmt.Println("Item termahal", data[IDXmahal].nama)
    fmt.Println("Item termurah", data[IDXmurah].nama)
}