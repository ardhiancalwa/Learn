package main
import "fmt"

func main()  {
	var pesan string
	
	fmt.Scan(&pesan)
	if len(pesan) >= 20 {
		fmt.Printf("%c%c%c%c",pesan[0],pesan[4],pesan[8],pesan[12])
	}
}