package main
import "fmt"

func main()  {
	var match1, match2, match3, match4, match5 string

	fmt.Scan(&match1, &match2, &match3, &match4, &match5)
	if match1 == "kalah" && match2 == "kalah" && match3 == "kalah" && match4 == "kalah" && match5 == "kalah" {
		fmt.Println("dipecat")
	} else {
		fmt.Println("tidak dipecat")
	}
}	