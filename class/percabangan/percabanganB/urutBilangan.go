package main
import "fmt"

func main()  {
	var bil1, bil2, bil3 int	

	fmt.Scan(&bil1, &bil2, &bil3)
	if bil1 <= bil2 && bil2 <= bil3 {
		fmt.Println(bil1, bil2, bil3)
	}
	if bil1 >= bil2 && bil2 >= bil3 {
		fmt.Println(bil3, bil2, bil1)
	}
	if bil1 >= bil2 && bil2 <= bil3 && bil1 <= bil3 {
		fmt.Println(bil2, bil1, bil3)
	}
}