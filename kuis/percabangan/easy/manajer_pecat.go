package main
import "fmt"

func main()  {
	var game1, game2, game3, game4, game5 string
	var pecat string

	fmt.Scan(&game1, &game2, &game3, &game4, &game5)
	if game1 == "kalah" && game2 == "kalah" && game3 == "kalah" && game4 == "kalah" && game5 == "kalah" {
		pecat = "dipecat"
	} else {
		pecat = "tidak dipecat"
	}
	fmt.Println(pecat)
}