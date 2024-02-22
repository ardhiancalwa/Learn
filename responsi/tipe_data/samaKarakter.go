package main
import "fmt"

func main()  {
	var char1, char2 byte
    var sama bool

    fmt.Scanf("%c %c", &char1, &char2)
    fmt.Println(char1, char2)
    sama =  (char1 == char2) || (char1 == (char2 - 32)) || (char1 == (char2 + 32)) && (char1 >= 65 && char1 <= 90) && (char2 >= 97 && char2 <= 122) || (char1 <= 122 && char1 >= 97) && (char2 <= 90 && char2 >= 65)
    fmt.Println(sama)
}