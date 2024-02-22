package main
import "fmt"

func main()  {
	var gigi int
	var kopling_ditarik, status_gas bool

	fmt.Scan(&gigi, &kopling_ditarik, &status_gas)
	if gigi == 0 {
		if kopling_ditarik && status_gas {
			fmt.Println("Mesin menyala dan motor tidak jalan")
		} else {
			fmt.Println("Mesin mati")
		}
	}

	if gigi == 1 || gigi == 2 || gigi == 3 || gigi == 4 {
		if kopling_ditarik && status_gas || !kopling_ditarik && status_gas {
			fmt.Println("Motor berjalan")
		} else if kopling_ditarik && !status_gas {
			fmt.Println("Mesin menyala dan motor tidak jalan")
		} else {
			fmt.Println("Mesin mati")
		}
	}
}