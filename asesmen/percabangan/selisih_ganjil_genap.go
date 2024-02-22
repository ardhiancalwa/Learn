package main
import "fmt"

func main() {
	var heiji, shinichi int
	var selisih int

	fmt.Scan(&heiji, &shinichi)
	selisih = heiji - shinichi
	if selisih == 0 {
		fmt.Println("Shinichi dan Heiji seri")
	} else if selisih%2 == 0 || selisih > 0 {
		if selisih > 0 && selisih<= 45 {
			fmt.Println("Shinichi menang")
		} else {
			fmt.Println("Shinichi dan Heiji kalah")
		}
	} else if selisih%2 != 0 || selisih < 0 {
		if selisih >= -16 && selisih < 0 {
			fmt.Println("Heiji menang")
		} else {
			fmt.Println("Shinichi dan Heiji kalah")
		}
	} else {
		fmt.Println("Shinichi dan Heiji kalah")
	}
}
