package main
import "fmt"

func main() {
	var bmi, tb, bb float64

	fmt.Scan(&tb, &bb)
	bmi = bb / (tb*tb)
	if bmi > 22.9 {
		fmt.Println("Kelebihan berat badan.")
	} else if bmi >= 18.5 {
		fmt.Println("Berat badan normal.")
	} else if bmi < 18.5 {
		fmt.Println("Berat badan kurang.")
	}
}