package main
import "fmt"

func main()  {
	var k rune
	var alphanumeric bool
	
	fmt.Scan(&k)
	alphanumeric = (k >= 'a' && k <= 'z') || (k >= 0 && k <= 9) || (k >= 'A' && k <= 'Z')
	fmt.Println(alphanumeric)
}