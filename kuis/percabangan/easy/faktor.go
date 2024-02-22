package main
import "fmt"

func main()  {
	var a, b int
	
	fmt.Scan(&a, &b)
	if b % a == 0 {
		fmt.Println("Ya,", a, "faktor", b)
	} else {
		fmt.Println("Tidak,", a, "bukan faktor", b)
	}
}