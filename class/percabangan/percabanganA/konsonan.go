package main
import "fmt"

func main()  {
	var karakter byte

	fmt.Scanf("%c", &karakter)
	if karakter != 'a' && karakter != 'i' && karakter != 'u' && karakter != 'e' && karakter != 'o' && karakter != 'A' && karakter != 'I' && karakter != 'U' && karakter != 'E' && karakter != 'O' {
		fmt.Println("konsonan")
	} else {
		fmt.Println("Bukan konsonan")
	}
}