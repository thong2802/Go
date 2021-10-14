package main

import "fmt"


func main() {
	// 1. if else
	number := 10
	if number == 10 {
		fmt.Println("number = 10")
	}

	//if condition {//code } else{//  code}
	if number < 10 {
		fmt.Println("number < 10")
	} else {
		fmt.Println("number = 10")
	}

	// if statment condition {// code}
	if  a := 100; a > 100 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("a = 10")
	}

	// 2. switch - case
	// else if- else if ....
	switch number {
	case 1:
		fmt.Println("Number = 1")
		break // co the dung break hoac khong dung
	case 2:
		fmt.Println("Number = 2")
	case 3:
		fmt.Println("Number = 3")
	case 4:
		fmt.Println("Number = 4")
	case 5:
		fmt.Println("Number = 5")
	case 6:
		fmt.Println("Number = 6")
	default:
		fmt.Println("unknow")
	}

	switch number {
	case 1, 200, 66:
		fmt.Println("Number = 1")
		break // co the dung break hoac khong dung
	case 2:
		fmt.Println("Number = 2")
	case 3:
		fmt.Println("Number = 3")
	case 4:
		fmt.Println("Number = 4")
	case 5:
		fmt.Println("Number = 5")
	case 6:
		fmt.Println("Number = 6")
	default:
		fmt.Println("unknow")
	}


	// FallThough -> giong continute
	switch number {
	case 1:
		fmt.Println("Number = 1")
		break // co the dung break hoac khong dung

	case 2:
		fmt.Println("Number = 2")
		fallthrough
	case 10:
		fmt.Println("Number = 3")
		fallthrough
	case 4:
		fmt.Println("Number = 4")
		fallthrough
	case 5:
		fmt.Println("Number = 5")
		fallthrough
	case 6:
		fmt.Println("Number = 6")
		fallthrough
	default:
		fmt.Println("unknow")
		
	}


	// break, goto
	switch number {
	case 1:
		fmt.Println("Number = 1")

	case 2:
		fmt.Println("Number = 2")
		fallthrough
	case 10:
		fmt.Println("Number = 3")
		break 
		fallthrough
	
	case 4:
		fmt.Println("Number = 4")
		fallthrough
	case 5:
		fmt.Println("Number = 5")
		fallthrough
	case 6:
		fmt.Println("Number = 6")
		fallthrough
	default:
		fmt.Println("unknow")
		
	}


	// for
	for i := 0; i < 10; i ++ {
		if i == 4 {
			continue
			//break
		}
		fmt.Println(i)
	}
	fmt.Println("out of Loop")


	// while -> go khong ho tro while tuy nhien ta co the mo phong no bang for

	j := 0
	for; j < 10;{
		fmt.Println(j)
		j ++
	}

	

	// infinite  loop-> vong lap vo tan trong go
	// for {
	// 	fmt.Println("9")
	// }

	//for init, condition, post
	for i, j := 1, 2; i < 10 && j < 10; i, j = i + 1, j + 1 {
		fmt.Println(i)
		fmt.Println(j)
	}




 }