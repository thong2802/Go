package main

import "fmt"

func main() {
//	check evenNumber, old
//	input a from keyboard
//	check a
	fmt.Println("Input a form key board")
	var a int
	fmt.Scan(&a)
	if a % 2 == 0{
		fmt.Println("Even Number")
	}else {
		fmt.Println("Old number")
	}


	for i := 0; i < 10; i++ {
		fmt.Println("Input a form key board")
		var a int
		fmt.Scan(&a)
		if a % 2 == 0{
			fmt.Println("Even Number")
		}else {
			fmt.Println("Old number")
		}
	}

	//for vo han
	for {
		fmt.Println("Input a form key board")
		var a int
		fmt.Scan(&a)
		if a % 2 == 0{
			fmt.Println("Even Number")
		}else {
			fmt.Println("Old number")
		}
	}

	//trong Go , ví dụ a = 5, a&1 = 1 => a lẽ, a = 4, a&1 = 0 => a chẵn
	fmt.Println("Input a form key board")
	var b int
	fmt.Scan(&a)
	if b & 1 ==0 {
		fmt.Println(b, "A is Even Number")
	}else {
		fmt.Println(b, "A is Old number")
	}
}
