package main

import "fmt"

func change(a *int)  {
	*a = 10
	fmt.Println("value a:" , a) // con trỏ a nằm ở một địa chỉ khác nhưng giá trị của nó lại lưu giữ địa chỉ bộ nhớ của biến khác
	fmt.Println("adddress a:" , &a)
}

func exchange(m *int, n *int)  {
	var temp = *m
		 *m	 = *n
		 *n	 = temp

	fmt.Println(m,n)

}

func changeArr(arr *[3]int)  {
	(*arr)[0] = 10
}
func main() {
	//Con trỏ là gì?
	//Một con trỏ là một biến chứa địa chỉ bộ nhớ của biến khác.
	var a int = 10
	fmt.Println(a)
	// get address in ram
	fmt.Println(&a)

	// pointer y
	var y *int = &a
	fmt.Println(y)
	fmt.Println(*y)

	// null -> java, c#, php
	// nil -> swift, golang

	// thay doi gia tri cua a
	*y = 6
	fmt.Println(a)


	fmt.Println("----------------")
	f := 7
	fmt.Println("address f:" , &f)

	change(&f)
	fmt.Println(f)

	//
	fmt.Println("========================")
	var m int = 5
	var n int = 10
	fmt.Println(m)
	fmt.Println(n)

	exchange(&m, &n)
	fmt.Println(m)
	fmt.Println(n)

	fmt.Println("========================")

	arr :=[3]int{4,5,6}
	changeArr(&arr)
	fmt.Println(arr)
	
}

