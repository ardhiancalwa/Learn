package main

import "fmt"

func main() {
	var tebak1, tebak2 int
    var noGenap_Ganjil int

    fmt.Scan(&tebak1, &tebak2)
    for (tebak1%2 == 0 && tebak2%2 != 0) || (tebak1%2 != 0 && tebak2%2 == 0) {
        if (tebak1%2 == 0 || tebak2%2 != 0) || (tebak1%2 != 0 && tebak2%2 == 0)  {
            noGenap_Ganjil++
        }
        fmt.Scan(&tebak1, &tebak2)
    }
    fmt.Println(noGenap_Ganjil)
}