package main
import "fmt"

func main()  {
	var u1, u2, u3 int
	var nilai string

	fmt.Scan(&u1, &u2, &u3)
	if u1 % 2 == 0 && u2 % 2 == 0 && u3 % 2 == 0 {
		nilai = "genap"
	} else if u1 % 2 != 0 && u2 % 2 != 0 && u3 % 2 != 0 {
		nilai = "ganjil"
	} else {
		nilai = "bukan ganjil atau genap"
	}
	fmt.Println(nilai)
}