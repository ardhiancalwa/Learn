package main
import "fmt"

func main() {
	var i, bilangan, bilanganAkhir, banyak int

	i = 0
	banyak = 0
	for {
		fmt.Scan(&bilangan)
		if i >= 10 || (i > 0 && bilangan > bilanganAkhir) {
			break
		}
		bilanganAkhir = bilangan
		banyak++
		i++
	}
	fmt.Println(banyak)
}
