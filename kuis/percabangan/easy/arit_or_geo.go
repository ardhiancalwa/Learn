package main
import "fmt"

func main()  {
	var u1, u2, u3 int
	var nilai string

	fmt.Scan(&u1, &u2, &u3)
	if u2-u1 == u3-u2 {
		nilai = "aritmatika"
	} else if u2/u1 == u3/u2 {
		nilai = "geometri"
	} else {
		nilai = "bukan aritmatika atau geometri"
	}
	fmt.Println(nilai)
}