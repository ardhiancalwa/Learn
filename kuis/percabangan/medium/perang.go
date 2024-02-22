package main
import "fmt"

func main() {
	var viking, saxon int
	
	fmt.Scan(&viking, &saxon)
	if viking * 4 > saxon {
		fmt.Println("viking menang")
	} else if viking * 4 < saxon {
		fmt.Println("saxon menang")
	} else {
		fmt.Println("imbang")
	}
}