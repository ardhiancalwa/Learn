package main
import "fmt"

func main()  {
	var x, nilai int

	fmt.Scan(&x)
	if x > 0 {
		nilai = x
	} else if x < 0 {
		nilai = x * -1
	} else {
		nilai = 0
	}
	fmt.Println(nilai)
}