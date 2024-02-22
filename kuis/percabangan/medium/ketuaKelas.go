package main
import "fmt"

func main()  {
	var jumSiswa, jumA, jumB int
	var minimal int

	
	fmt.Scan(&jumSiswa, &jumA, &jumB)
	minimal = jumSiswa - (jumSiswa * 60/100)
	fmt.Println(minimal)
	if minimal <= (jumA+jumB) && jumA > jumB {
		fmt.Println("Kandidat A menang")
	} else if minimal <= (jumA+jumB) && jumB > jumA{
		fmt.Println("Kandidat B menang")
	} else {
		fmt.Println("Tidak ada pemenang")
	}
}