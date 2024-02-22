package main
import "fmt"

func main() {
	var kota1, kota2 string

	fmt.Scan(&kota1, &kota2)
	if kota1 == "bogor" && kota2 == "sukabumi" || kota1 == "sukabumi" && kota2 == "bogor" && kota1 != kota2 {
		fmt.Println("bertetanggaan")
	} else if kota1 == "bogor" && kota2 == "cianjur" || kota1 == "cianjur" && kota2 == "bogor" && kota1 != kota2 {
		fmt.Println("bertetanggaan")
	} else if kota1 == "sukabumi" && kota2 == "cianjur" || kota1 == "cianjur" && kota2 == "sukabumi" && kota1 != kota2 {
		fmt.Println("bertetanggaan")
	} else if kota1 == "cianjur" && kota2 == "bandung" || kota1 == "bandung" && kota2 == "cianjur" && kota1 != kota2 {
		fmt.Println("bertetanggaan")
	} else if kota1 == "bandung" && kota2 == "tasikmalaya" || kota1 == "tasikmalaya" && kota2 == "bandung" && kota1 != kota2 {
		fmt.Println("bertetanggaan")
	} else {
		fmt.Println("tidak bertetanggaan")
	}
}