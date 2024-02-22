package main
import "fmt"

func main()  {
	var a, b int

	fmt.Scan(&a, &b)
	if a % b == 0 {
		fmt.Println("Ya,", a, "kelipatan", b)
	} else {
		fmt.Println("Tidak,", a, "bukan kelipatan", b)
	}
}