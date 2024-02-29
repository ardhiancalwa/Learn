package main

import "fmt"

func segitiga(s1, s2, s3 int) string {
	if s1+s2 > s3 && s1+s3 > s2 && s2+s3 > s1 {
		return "segitiga"
	} else {
		return "bukan segitiga"
	}

}

func main() {
	var s1, s2, s3 int
	fmt.Scan(&s1, &s2, &s3)
	fmt.Println(segitiga(s1, s2, s3))
}
