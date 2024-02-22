package main

import (
	"fmt"
	// "math"
)

var nama string

func main() {
	// sayHello()
	// pangkat()
	// variable()
	// constanta()
	// InO()
	DivMod()
}

// func sayHello() {
// 	fmt.Println("Hello World")

// 	nama = "Calwaa"
// 	fmt.Println("Namaku", nama)

// 	var fullName string = "Ardhian Calwa Nugraha"
// 	fmt.Println("Nama Lengkap", fullName)

// 	var name1 string
// 	var umur int
// 	fmt.Print("Nama Saya: ")
// 	fmt.Scan(&name1)
// 	fmt.Print("Umur saya: ")
// 	fmt.Scan(&umur)
// 	fmt.Println("Ohh nama kamu: ", name1)
// 	fmt.Println("Ohh umur kamu: ", umur)
// }

// func pangkat() {
// 	fmt.Println(math.Pow(3.0, 2.0)) //math.pow digunakan untuk perpangkatan
// 	fmt.Println(1 + 2)
// 	fmt.Println(1 - 2)
// 	fmt.Println(1 * 2)
// }

// func variable() {
// 	var bil_bulat int
// 	var bil_pecahan float64
// 	var karakter byte // char, ada byte dan rune
// 	var ketemu bool
// 	var nama string

// 	bil_bulat = 3
// 	bil_pecahan = 1.5
// 	karakter = 'c'
// 	ketemu = true
// 	nama = "ronaldo"

// 	fmt.Println(bil_bulat)
// 	fmt.Println(bil_pecahan)
// 	fmt.Println("%c\n", karakter)
// 	fmt.Println(ketemu)
// 	fmt.Println(nama)
// }

// func constanta()  {
// 	const pi float64 = 3.14
// 	var r float64

// 	r = 10.0

// 	// fmt.Pow()
// 	fmt.Println(pi * math.Pow(r,2))
// }

// func InO()  {
// 	var karakter byte	

// 	fmt.Scan(karakter)
// 	fmt.Println(karakter)
// }

func DivMod()  {
	var bil, d1, d2, d3 int

	fmt.Scan(&bil)
	d1 = bil / 100
	d2 = bil / 10 % 10
	d3 = bil % 10

	fmt.Println(d1, d2, d3)
}