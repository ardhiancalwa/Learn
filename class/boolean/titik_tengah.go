package main
import "fmt"

func main()  {
	var bil1, bil2, bil3 int
	var valid bool

	fmt.Scan(&bil1, &bil2, &bil3)
	valid = ((bil1 + bil2) / 2 == bil3) || ((bil2 + bil3) / 2 == bil1) || ((bil1 + bil3) / 2 == bil2) && (bil1 != 0 && bil2 != 0 && bil3 != 0)
	fmt.Println(valid)
}