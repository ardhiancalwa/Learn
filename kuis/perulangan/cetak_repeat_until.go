package main
import "fmt"

func main() {
    var i, n, m int
    fmt.Scan(&n, &m)
    i = n
    for {
        fmt.Println(i)
        if i >= m {
            break
        }
        i++
    }
}