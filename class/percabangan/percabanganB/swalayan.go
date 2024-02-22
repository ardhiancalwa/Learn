package main
import "fmt"

func main() {
	var cashback, diskon float64
	var membership string
	var a, b, c, d, e int
	const hitungCashback float64 = 3.1
	const hitungDiskon float64 = 1.7

	fmt.Scan(&membership, &a, &b, &c, &d, &e)
	if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		if membership == "Yes" {
			cashback = float64(a+b+c) * (hitungCashback) + 15
			diskon = 0
			if cashback > 35 {
				cashback = 35
			}
			fmt.Println(membership, "Cashback:",cashback,"%", "Diskon:", diskon,"%")
		} else {
			cashback = float64(a+b+c) * (hitungCashback)
			diskon = 0
			if cashback > 35 {
				cashback = 35
			}
			fmt.Println(membership, "Cashback:",cashback,"%", "Diskon:", diskon,"%")
		}
	} else if a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0 {
		if membership == "Yes" {
			diskon = float64(c+d+e) * hitungDiskon + 15
			cashback = 0
			if diskon > 50 {
				diskon = 50
			}
			fmt.Println(membership, "Cashback:", cashback, "%", "Diskon", diskon,"%")
		} else	{
			diskon = float64(c+d+e) * hitungDiskon
			cashback = 0
			if diskon > 50 {
				diskon = 50
			}
			fmt.Println(membership, "Cashback:", cashback,"%", "Diskon", diskon,"%")
		}
	} else if a%2 == 0 || b%2 == 0 || c%2 == 0 || d%2 == 0 || e%2 == 0 {
		if membership == "Yes" {
			cashback = float64(a+b+c) * (hitungCashback) + 15
			diskon = float64(c+d+e) * hitungDiskon + 15
			if cashback > 35 {
				cashback = 35
			}
			if diskon > 50 {
				diskon = 50
			}
			fmt.Println(membership, "Cashback:", cashback,"%", "Diskon:", diskon,"%")
		} else {
			cashback = float64(a+b+c) * (hitungCashback)
			diskon = float64(c+d+e) * hitungDiskon
			if cashback > 35 {
				cashback = 35
			}
			if diskon > 50 {
				diskon = 50
			}
			fmt.Println(membership, "Cashback:", cashback,"%", "Diskon:", diskon,"%")
		}
	}
}
