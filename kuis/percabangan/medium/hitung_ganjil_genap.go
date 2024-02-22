package main
import "fmt"

func main() {
	var bil1, bil2, bil3, bil4 int
	var ganjil, genap int

	fmt.Scan(&bil1, &bil2, &bil3, &bil4)
	if bil1%2 == 0 {
		genap++
	} else {
		ganjil++
	}
	if bil2%2 == 0 {
		genap++
	} else {
		ganjil++
	}
	if bil3%2 == 0 {
		genap++
	} else {
		ganjil++
	}
	if bil4%2 == 0 {
		genap++
	} else {
		ganjil++
	}

	fmt.Println(genap, ganjil)
}
