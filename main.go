package main

import (
	"fmt"
	"math"
)


func main()  {
	fmt.Println("Hello Golang toturial")
	fmt.Println(Sum(3))

	fmt.Println("--------------------------------------------------")

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
	fmt.Println("--------------------------------------------------")
}
func Sum(number int) (int) {
	sum:=0
	for i := 0; i < number; i++ {
		sum += i
	}
	return sum
}