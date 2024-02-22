package main
import "fmt"

func main()  {
	var karakter byte
	for karakter != '.' {
		fmt.Printf("%c", karakter)
		fmt.Scanf("%c", &karakter)
	}
}