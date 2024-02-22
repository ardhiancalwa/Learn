package main
import "fmt"

func main()  {
	var bilangan int
	var hari string

	fmt.Scan(&bilangan)
	if bilangan == 1 {
		hari = "minggu"
	} else if bilangan == 2 {
		hari = "senin"
	} else if bilangan == 3 {
		hari = "selasa"
	} else if bilangan == 4 {
		hari = "rabu"
	} else if bilangan == 5 {
		hari = "kamis"
	} else if bilangan == 6 {
		hari = "jumat"
	} else if bilangan == 7 {
		hari = "sabtu"
	}
	fmt.Println(hari)
}