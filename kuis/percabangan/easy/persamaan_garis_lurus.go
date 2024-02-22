package main
import "fmt"

func main()  {
	var m, c float64
	var titik_asal string

	fmt.Scan(&m, &c)
	if c == 0 {
		titik_asal = "melewati"
	} else {
		titik_asal = "tidak melewati"
	}
	fmt.Println(titik_asal)
}