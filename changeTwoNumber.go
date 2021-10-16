package main

import "fmt"

func Swap( a int,  b int) {
	fmt.Println("a: ",  a)
	fmt.Println("b: ", b)
	temp := a
	a = b
	b = temp
	fmt.Println("a: ",  a)
	fmt.Println("b: ", b)
}
func main() {
	// user tempt avairiable
	// a = 4
	// b = 5

	// c = a
	// a = b
	// b = c
	Swap(4,5)


	// khong su dung bien tam
	var c int
	var d int
	fmt.Println("Input c")
	fmt.Scan(&c)
	fmt.Println("Input d")
	fmt.Scan(&d)

	fmt.Println("c: ",  c)
	fmt.Println("d: ", d)

	d = c + d
	c = d - c
	fmt.Println("c: ",  c)
	d = d - c
	fmt.Println("d: ", d)

	// su dung pointer
	

}
