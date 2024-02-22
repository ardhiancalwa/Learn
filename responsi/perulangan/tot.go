package main
import "fmt"

func main(){
	var answer string
	var score, count int
	var check bool
	
	fmt.Scan(&answer)
	score = 0
	count = 0
	check = (answer != "TTT" && ((answer == "OOO")||(answer == "OOT")||(answer == "OTO")||(answer == "TOO")||(answer == "OOT")||(answer == "TOT")||(answer == "TTO")))
	for check == true && count != 10 {
		count += 1
		score += 1
		fmt.Scan(&answer)
		check = (answer != "TTT" && ((answer == "OOO")||(answer == "OOT")||(answer == "OTO")||(answer == "TOO")||(answer == "OOT")||(answer == "TOT")||(answer == "TTO")))
	}
	fmt.Println("Poin:",score)
}