package main
import "fmt"

func main()  {
    var bilangan1, bilangan2, bilanganTukar1, bilanganTukar2, digitRatusan, digitSatuan int
    fmt.Scan(&bilangan1, &bilangan2)
    
    digitRatusan = bilangan1 / 100
    digitSatuan = bilangan2 % 10
    bilanganTukar1 = digitSatuan*100 + bilangan1 % 100
    bilanganTukar2 = (bilangan2 / 10 * 10) + digitRatusan
    fmt.Println(bilanganTukar1, bilanganTukar2)
}