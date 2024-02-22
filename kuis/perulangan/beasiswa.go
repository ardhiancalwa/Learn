package main
import "fmt"

func main()  {
	var beasiswa string
	
	fmt.Scan(&beasiswa)
	for beasiswa != "ada" || beasiswa == "belum" {
		fmt.Print("cari beasiswa")
		fmt.Scan(&beasiswa)
	}
	fmt.Print("pencarian selesai")
}