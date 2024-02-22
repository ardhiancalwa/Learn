package main
import "fmt"

func main()  {
    var upperCase, lowerCase rune
    fmt.Scanf("%c", &upperCase)
    if upperCase >= 'A' && upperCase <= 'Z' {
        lowerCase = upperCase + 32
    }
    fmt.Printf("%c", lowerCase)
}