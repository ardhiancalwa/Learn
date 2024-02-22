package main
import "fmt"

func main()  {
	var AB, BC, CD, AD int

	fmt.Scan(&AB, &BC, &CD, &AD)
	if AB > 0 && BC > 0 && CD > 0 && AD > 0 {
		if AB == CD && BC == AD {
			fmt.Println("Anda dapat membentuk persegi panjang")
		} else if AB == BC && BC == CD && CD == AD {
			fmt.Println("Anda dapat membentuk persegi")
		} else {
			fmt.Println("Anda tidak membentuk bidang apapun")
		}
	} else {
		fmt.Println("Anda harus lebih dari 0")
	}
}