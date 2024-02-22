package main
import "fmt"

func main() {
    var i, jumlah int

	jumlah = 0
	for i = 1; i <= 100; i++ {
		jumlah += i
	}
	fmt.Println(jumlah)
}
