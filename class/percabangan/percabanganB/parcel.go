package main
import "fmt"

func main()  {
	var berat, total int
	var kg, gram int
	const biayaKirim int = 10000

	fmt.Scan(&berat)
	kg = berat / 1000
	gram = berat % 1000
	fmt.Println(kg, "kg", "+", gram, "gram")

	if kg < 10 && gram >= 500 {
		kg *= biayaKirim
		gram *= 5
		total = kg + gram
		fmt.Println("Rp.", kg, "+", gram, "gram", total)
	} else if kg < 10 && gram < 500 {
		kg *= biayaKirim
		gram *= 15
		total = kg + gram
		fmt.Println("Rp.", kg, "+", gram, "gram", total)
	} else {
		kg *= biayaKirim
		gram = 0
		total = kg + gram
		fmt.Println("Rp.", kg, "+", gram, "gram", total)
	} 
}