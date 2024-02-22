package main
import "fmt"

func main()  {
	var segitiga bool
	var s1, s2, s3 int

	fmt.Scan(&s1, &s2, &s3)
	segitiga = (s1 + s2 > s3) && (s2 + s3 > s1) && (s1 + s3 > s2)
	fmt.Println(segitiga)
}