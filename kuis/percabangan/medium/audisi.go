package main
import "fmt"

func main()  {
	var vote1, vote2, vote3 string

	fmt.Scan(&vote1, &vote2, &vote3)
	if vote1 == "yes" && vote2 == "yes" && vote3 == "yes" {
		fmt.Println("lolos")
	} else if (vote1 == "yes" && vote2 == "yes" && vote3 == "no") || (vote1 == "yes" && vote2 == "no" && vote3 == "yes") || (vote1 == "no" && vote2 == "yes" && vote3 == "yes") {
		fmt.Println("lolos")
	} else {
		fmt.Println("tidak lolos")
	}
}