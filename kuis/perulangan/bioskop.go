	package main
	import "fmt"

	func main()  {
		var i, jam, menit, jumJam, jumMenit int

		fmt.Scan(&jam, &menit)
		i = 0
		for jam != 0 && menit != 0 {
			jumJam += jam
			jumMenit += menit 
			i++
			fmt.Scan(&jam, &menit)
		}
		fmt.Println(jumJam, jumMenit)
	}