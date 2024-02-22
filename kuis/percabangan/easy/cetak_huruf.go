package main
import "fmt"

func main()  {
	var huruf string
	var deskripsi int

	fmt.Scan(&huruf)
	if huruf == "A" {
		deskripsi = 5
		fmt.Println(deskripsi)
	} else if huruf == "B" {
		deskripsi = 4
		fmt.Println(deskripsi)
	} else if huruf == "C" {
		deskripsi = 3
		fmt.Println(deskripsi)
	} else if huruf == "D" {
		deskripsi = 2
		fmt.Println(deskripsi)
	} else if huruf == "E" {
		deskripsi = 1
		fmt.Println(deskripsi)
	} else if huruf == "T" {
		deskripsi = 0
		fmt.Println(deskripsi)
	}
}