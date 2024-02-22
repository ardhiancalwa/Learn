	package main
	import "fmt"

	func main()  {
		var saldo, transaksi int

		saldo = 0
		fmt.Scan(&transaksi)
		for transaksi != 0 {
			saldo += transaksi
			fmt.Scan(&transaksi)
		}
		fmt.Println(saldo)
	}