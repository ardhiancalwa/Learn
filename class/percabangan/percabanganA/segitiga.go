package main
import "fmt"

func main()  {
	var s1, s2, s3 int
	var segitiga string	

	fmt.Scan(&s1, &s2, &s3)
	if s1 == s2 && s1 == s3 && s2 == s3 {
		segitiga = "Segitiga Sama Sisi"
	} else if s1 + s2 > s3 && s1 + s3 > s2 && s2 + s3 > s1 {
		segitiga = "Segitiga Sama Kaki"
	} else if s1 < s2 && s2 < s3 && s1 < s3 {
		segitiga = "Segitiga Sembarang"
	}
	fmt.Println(segitiga)
}