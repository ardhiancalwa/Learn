package main
import "fmt"

func main()  {
	var rangking int
	var hadiah string

	fmt.Scan(&rangking)
	if rangking <= 5 {
		hadiah = "mendapat hadiah"
	} else {
		hadiah = " "
	}
	fmt.Println(hadiah)
}