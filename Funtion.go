package main

import "fmt"

// func func_name (params) return type {...}

func Chao() {
	fmt.Println("Hi go lang")
}

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func greeting (name string) string {
	result := fmt.Sprintf("Hello %s", name)
	return result
}


// mutiple return value
func recInfo(w int, h int) (int, int) {
	return w, h
}

// Named return value
func recInfo1(w int, h int) (width int,heigth int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}

func main() {
	Chao()
	SayHello("Thong")
	result := greeting("Thong")
	fmt.Println(result)
	w, h := recInfo(100, 200)
	fmt.Println("Width:= ", w)
	fmt.Println("Height:= ", h)


	/////////////////
	w, h, isSquare := recInfo1(200, 200)
	if isSquare {
		fmt.Println("this is Square")
	}else {
		fmt.Println("Width: ", w)
		fmt.Println("heigth: ", h)
	}

	

}




