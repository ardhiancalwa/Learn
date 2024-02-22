package main
import "fmt"

func main() {
	var bil1, bil2, bil3, bil4 int
	var negatif, positif int

	fmt.Scan(&bil1, &bil2, &bil3, &bil4)
	if bil1 > 0 {
		positif++
	} else if bil1 < 0 {
		negatif++
	}
	if bil2 > 0 {
		positif++
	} else if bil2 < 0 {
		negatif++
	}
	if bil3 > 0 {
		positif++
	} else if bil3 < 0 {
		negatif++
	}
	if bil4 > 0 {
		positif++
	} else if bil4 < 0 {
		negatif++
	}

	fmt.Println(positif, negatif)
}
