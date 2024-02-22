package main
import "fmt"

func main() {
	var jam, menit, tarif, waktuPesan int
	var harga, jarak float64
	const waktuBuka int = 5 * 60
	const waktuTutup int = 22 * 60

	fmt.Scan(&jam, &menit, &jarak)
	jam *= 60
	waktuPesan = jam + menit
	if waktuPesan >= waktuBuka && waktuPesan <= waktuTutup {
		if (waktuPesan >= waktuBuka && waktuPesan <= 9*60) || (waktuPesan >= 16*60 && waktuPesan <= 19*60) {
			if jarak > 0 && jarak <= 10 {
				tarif = 5000
				harga = jarak * float64(tarif)
			} else if jarak > 10 && jarak <= 20 {
				tarif = 4500
				harga = jarak * float64(tarif)
			}
		} else {
			if jarak > 0 && jarak <= 20 {
				tarif = 4000
				harga = jarak * float64(tarif)
			}
		}
		fmt.Println(harga)
	} else {
		fmt.Println("Maaf, kami tidak bisa melayani pesanan anda.")
	}
}
