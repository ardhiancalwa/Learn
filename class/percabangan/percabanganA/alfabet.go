package main
import "fmt"

func main()  {
	var karakter byte

	fmt.Scanf("%c", &karakter)
	if karakter >= 'a' && karakter <= 'z' || karakter >= 'A' && karakter <= 'Z' {
		fmt.Println("Alfabet")
	} else {
		fmt.Println("Bukan alfabet")
	}
}