package main
import "fmt"

func main()  {
	// var i int

	// for i <= 5 {
	// 	fmt.Println("hello")
	// 	i += 1
	// }

	var i int
	var selesai bool

	i = 1
	selesai = false
	for i <= 10 && !selesai {
		fmt.Println(i)
		selesai = i % 4 == 0
		i++
	}
	
}