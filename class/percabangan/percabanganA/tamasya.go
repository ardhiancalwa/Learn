package main
import "fmt"

func main()  {
	var N, mobil int

	fmt.Scan(&N)
	mobil = N / 7
	if N % 7 != 0 {
		mobil++
	}
	fmt.Println(mobil)
}