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
	var myByte byte = 255
	fmt.Println(myByte)
	fmt.Printf("\n%T", myByte)
	fmt.Println(math.MaxUint8)

	// byte // alias for uint8 alias for uint8
	var a byte = 'E'
	fmt.Println(a)
	fmt.Printf("%X", a)
	
	fmt.Println("---------------")
	// float32, complex 64 (so thuc)
	var b float64 = 5.9
	fmt.Println(b)
	
	// complex 64 complex 128 ( so phuc )
    var z1 complex64 = 10 + 1i
	fmt.Println(z1)

	var z2 complex64 = 20 + 1i
	fmt.Println(z1 + z2)


	////////////////////////////////////////////
	// Rune -> int32
	// khong quan tam ki tu do co bao nhieu byte ma ki tu do co bao nhieu byte thi no se cap phat bo nho cho no
	var String string = "Nhẫn"
	fmt.Println(String)

	for i := 0; i < len(String); i++ {
		fmt.Printf("%c", String[i])
	}

	fmt.Println("\n-----------------")
    // ep kieu sang runes
	var String1 string = "Nhẫn"
	runes := []rune(String1)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}

	//
	var myRunes = 'A'
	fmt.Printf("\n%T", myRunes)

	// Type Convertion - chuyen doi kieu du lieu
	// Go rat nghiem ngat trong viec chuyen doi kieu du lieu
	var A int = 1
	var B float64 = 2.1
	fmt.Println(float64(A) + B)


	// unitptr -> luu giua dia chi con tro

}
