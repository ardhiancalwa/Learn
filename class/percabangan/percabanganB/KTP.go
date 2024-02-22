package main
import "fmt"

func main()  {
	var usia int
	var kk bool

	fmt.Scan(&usia, &kk)
	if usia >= 17 && kk {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}