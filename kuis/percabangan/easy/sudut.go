package main
import "fmt"

func main()  {
	var sudut1, sudut2, sudut3 int
	var segitiga bool

	fmt.Scan(&sudut1, &sudut2, &sudut3)
	if sudut1 + sudut2 + sudut3 == 180 {
		segitiga = true
	} else {
		segitiga = false
	}
	fmt.Println(segitiga)
}