package main
import "fmt"

func main()  {
	var temp int

	fmt.Scan(&temp)
	if temp < 0 {
		fmt.Println("Freezing weather")
	} else if temp < 9 {
		fmt.Println("Very Cold weather")
	} else if temp < 19 {
		fmt.Println("Cold weather")
	} else if temp < 29 {
		fmt.Println("Normal in Temp")
	} else if temp < 39 {
		fmt.Println("It's Hot")
	} else {
		fmt.Println("It's Very Hot")
	}
}