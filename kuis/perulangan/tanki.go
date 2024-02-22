package main
import "fmt"

func main()  {
	var kapasitas, volume_air, jumlahVolume int
	var penuh bool

	fmt.Scan(&kapasitas)
	jumlahVolume = 0
	for jumlahVolume < kapasitas {
		fmt.Scan(&volume_air)
		jumlahVolume += volume_air
		penuh = jumlahVolume == kapasitas || jumlahVolume >= kapasitas
		fmt.Println(penuh)
	}
}