package main
import "fmt"

func main()  {
	var n int
	var username, password string

	fmt.Scan(&username, &password)
	n = 0
	for username != "admin" || password != "admin"{
		n++
		fmt.Scan(&username, &password)
	}
	fmt.Println(n, "Login berhasil")
}