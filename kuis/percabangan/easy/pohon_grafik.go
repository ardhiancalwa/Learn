package main
import "fmt"

func main()  {
	var pohon string

	fmt.Scan(&pohon)
	if pohon == "A" {
		fmt.Println("akar")
	} else if pohon == "B" {
		fmt.Println("verteks dalam")
	} else if pohon == "C" {
		fmt.Println("daun")
	} else if pohon == "D" {
		fmt.Println("daun")
	} else if pohon == "E" {
		fmt.Println("verteks dalam")
	} else if pohon == "F" {
		fmt.Println("daun")
	} else if pohon == "G" {
		fmt.Println("daun")
	}
}