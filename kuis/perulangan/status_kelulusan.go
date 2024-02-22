package main
import "fmt"

func main() {
	var clo1, clo2, clo3 float64
	var lulus bool

	fmt.Scan(&clo1, &clo2, &clo3)
	for clo1 > 0 && clo2 > 0 && clo3 > 0 {
		lulus = true
		lulus = lulus && clo1 > 50 && clo2 > 50 && clo3 > 50
		fmt.Println(lulus)
		fmt.Scan(&clo1, &clo2, &clo3)
	}
}
