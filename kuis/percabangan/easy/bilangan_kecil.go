package main
import "fmt"

func main() {
	var bil1, bil2 int
	
	fmt.Scan(&bil1, &bil2)
	if bil1 < bil2 {
		fmt.Println(bil1)
	} else {
		fmt.Println(bil2)
	}
}