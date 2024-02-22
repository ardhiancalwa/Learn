package main

import "fmt"

func main()  {
    var s1, s2, s3 int
    var segitiga_sah bool
    fmt.Scan(&s1, &s2, &s3)

    segitiga_sah = s1 + s3 > s2 && s2 + s3 > s1 && s1 != 0 && s2!=0 && s3 != 0
    fmt.Println(segitiga_sah)
}