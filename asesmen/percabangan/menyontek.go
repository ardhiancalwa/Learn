package main
import "fmt"

func main()  {
	var nilai int
	var menyontek bool

	fmt.Scan(&nilai, &menyontek)
	if menyontek {
		nilai = 0
		fmt.Println(nilai)
	} else {
		fmt.Println(nilai)
	}
}