package main
import "fmt"

func main()  {
	var drachma int
	var oboli, mine, attic int

	fmt.Scan(&drachma)
	attic = drachma / 6000
	drachma = drachma % 6000
	mine = drachma / 100
	drachma = drachma % 100
	oboli = drachma * 6

	fmt.Print(attic, mine, oboli)
}