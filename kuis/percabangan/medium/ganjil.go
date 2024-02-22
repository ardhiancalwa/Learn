package main
import "fmt"

func main()  {
	var bil1, bil2, bil3, bil4, bil5 int

	fmt.Scan(&bil1, &bil2, &bil3, &bil4, &bil5)
	if bil1%2 != 0 && bil2%2 != 0 && bil3%2 != 0 && bil4%2 != 0 && bil5%2 != 0 {
		fmt.Println("ganjil semua")
	} else if bil1%2 != 0 || bil2%2 != 0 || bil3%2 != 0 || bil4%2 != 0 || bil5%2 != 0 {
		fmt.Println("ganjil sebagian")
	} else {
		fmt.Println("tidak ada bilangan ganjil")
	}
}