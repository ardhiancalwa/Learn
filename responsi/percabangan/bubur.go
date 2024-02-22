package main
import "fmt"

func main()  {
	var toping1, toping2, toping3, toping4 bool
	var harga int

	fmt.Scan(&toping1, &toping2, &toping3, &toping4)
	harga = 6000
	if toping1 && toping2 && toping3 && toping4 {
		harga = harga + 3000 + 1500 + 4500 + 4000
	}else if toping1 {
		harga += 3000
	} else if toping2 {
		harga += 1500
	} else if toping3 {
		harga += 4500
	} else if toping4 {
		harga += 4000
	}
	fmt.Println(harga)
}