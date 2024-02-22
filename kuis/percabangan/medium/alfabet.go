package main
import "fmt"

func main()  {
	var huruf byte

	fmt.Scanf("%c",&huruf)
	if huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' || huruf == 'A' || huruf == 'I' || huruf == 'U' || huruf == 'E' || huruf == 'O' {
		fmt.Println("vokal")
	} else {
		fmt.Println("konsonan")
	}
}