package main
import "fmt"

func main()  {
	var X int
	var konsekutif bool

	fmt.Scan(&X)
	konsekutif = true
	for X >= 10 {
		konsekutif = konsekutif && ((X % 10) - (X / 10 % 10) == 1) || ((X / 10 % 10) - (X % 10) == 1) 
		X /= 10
	}
	fmt.Println(konsekutif)
}