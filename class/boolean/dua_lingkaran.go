package main
import "fmt"

func main()  {
	var rA, rB, jarak int
	var irisan bool

	fmt.Scan(&rA, &rB, jarak)
	irisan = rA + rB > jarak
	fmt.Println(!irisan)
}