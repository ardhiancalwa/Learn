package main
import "fmt"

func main() {
    var digit1, digit2, digit3, digit4, nomorVoucher int
	var diskon, cashback, voucher bool
	fmt.Scan(&nomorVoucher)

	digit1 = nomorVoucher / 1000
	digit2 = (nomorVoucher / 10) % 100
	digit3 = (nomorVoucher / 10) % 10
	digit4 = nomorVoucher % 10 

	diskon = digit2 % 2 == 0
	cashback = (digit1 + digit3) % digit4 == 0
	voucher = (diskon || cashback) && !(diskon && cashback)
	fmt.Println(diskon, cashback, voucher)
}