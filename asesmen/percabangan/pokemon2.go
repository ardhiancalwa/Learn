package main
import "fmt"

func main()  {
	var rangee int

	fmt.Scan(&rangee)
	if rangee == 31 {
		fmt.Println("Best")
	} else if rangee == 30 {
		fmt.Println("Fantastic")
	} else if rangee >= 26 && rangee <= 29{
		fmt.Println("Very Good")
	} else if rangee >= 16 && rangee <= 25 {
		fmt.Println("Pretty Good")
	} else if rangee >= 1 && rangee <= 15{
		fmt.Println("Decent")
	} else if rangee == 0 {
		fmt.Println("No Good")
	}
}
