package main
import "fmt"

func main()  {
    var i, n, keuntungan, pendapatan, aset int

    fmt.Scan(&n)
    i = 1
    aset = 0
    for i <= n {
        fmt.Scan(&keuntungan)
        pendapatan += keuntungan / 3
        aset += keuntungan % 3
        i++
    }
    fmt.Println(pendapatan, aset)
}
