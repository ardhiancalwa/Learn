package main
import "fmt"

func main() {
	var kapasitas, jumPenumpang int
	var kapasitasMin, kapasitasMax float64

	fmt.Scan(&kapasitas, &jumPenumpang)
	kapasitasMin = 0.5 * float64(kapasitas)
	kapasitasMax = 0.75 * float64(kapasitas)

	if jumPenumpang >= int(kapasitasMin) && jumPenumpang <= int(kapasitasMax) {
		fmt.Println("berangkat")
	} else {
		fmt.Println("tidak berangkat")
	}
}
