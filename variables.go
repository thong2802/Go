package main

import "fmt"

func main ()  {
	var number int
	number = 10
	fmt.Println(number)

	var number1 int = 9
	fmt.Println(number1)
	
	
	// Type Reference
	var email = "ducthong2802@gmail.com"
	fmt.Println(email)

	// khai bao nhieu bien
	// 1. khai bao nhieu bien cung kieu du lieu
	var a, b int
	a = 1
	b = 4
	fmt.Println(a)
	fmt.Println(b)

	var a1, b1 int = 4, 5
	fmt.Println(a1)
	fmt.Println(b1)

	// Type Reference
	var a2, b2  = 6, 7
	fmt.Println(a2)
	fmt.Println(b2)

	// 2. khai bao nhieu bien khac kieu du lieu
    var (
		name string
		address string
		age int
	)
	name = "Nguyen duc thong"
	address = "TPHCM"
	age = 20

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(age)

	var (
		name1 string = "Thong"
		address1 string = "Singapore"
		age1 int = 20
	)

	fmt.Println(name1)
	fmt.Println(address1)
	fmt.Println(age1)

	// Type Reference
	var name2, address2, age2 = "Thong Nguyen", "Dak Lak", 20

	fmt.Println(name2)
	fmt.Println(address2)
	fmt.Println(age2)

	// short variable
	language := "go lang"
	
	fmt.Println(language)

	//  Constants
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

}