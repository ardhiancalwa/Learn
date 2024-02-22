package main
import "fmt"

func main()  {
	var bil1, bil2 int

	fmt.Scan(&bil1, &bil2)
	if bil1%2 == 0 && bil2%2 == 0 {
		fmt.Println(bil1*bil2)
	} else if bil1%2 != 0 && bil2%2 != 0 {
		fmt.Println(bil1+bil2)
	} else {
		fmt.Println(0)
	}
}