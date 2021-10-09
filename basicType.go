package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
   	// Boolean
	var myBool bool = true // fasle
	fmt.Println(myBool)

	var mySecondBool bool = false
	fmt.Println(mySecondBool)

	// string
	var MyString string  = 	"Go lang"
	fmt.Println(MyString)

	// int
	var MyInt int  = 123
	fmt.Println(MyInt)

	// int 8, 16, 32, 64 => so bit toi da ma no dai dien
	// Đối với kiểu dữ liệu kiểu int thì khi các bạn khai báo int trong chương trình thì tuỳ vào hệ điều hành (32 bit or 64 bit) mà go compiler sẽ tự động convert qua .
	// 1.Range
	// 2.Bits

	// 1. Range
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)

	fmt.Println("-----------")
	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)

	fmt.Println("-----------")
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)

	fmt.Println("-----------")
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	// 2.Bits
	fmt.Println("-----------")
	fmt.Println(bits.OnesCount8(math.MaxUint8))
	fmt.Println("-----------")
	fmt.Println(bits.OnesCount16(math.MaxUint16))
	fmt.Println("-----------")
	fmt.Println(bits.OnesCount32(math.MaxUint32))
	fmt.Println("-----------")
	fmt.Println(bits.OnesCount64(math.MaxUint64))


	//Uint -> số nguyên không dấu -> số nguyên dương
	var myUint uint = 10
	fmt.Println(myUint)

	// byte -> bản chất là kiểu dữ liệu Unit 8
	var myByte = 1
	fmt.Println(myByte)
	fmt.Printf("%T", myByte)









	

}
