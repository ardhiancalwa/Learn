package main
import "fmt"

func main()  {
	var bil1, bil2, bil3 int
	var konsekutif string

	fmt.Scan(&bil1, &bil2, &bil3)
	if bil3 - bil2 == 1 || bil2 - bil1 == 1 {
		konsekutif = "ya"
	} else {
		konsekutif = "tidak"
	}
	fmt.Println(konsekutif)
}