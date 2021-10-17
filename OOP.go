package main

import "fmt"

type Cat struct {
	id int
	name string
	age int
	leg int
	tail bool
}

type Rabbit struct {
	id int
	name string
	age int
	leg int
	tail bool
}

// trong struct khong ho tro viet ham ben trong nó nen ta phai viét ngoài rồi dùng con trỏ trỏ tới nó
func (c *Cat)eat()  {
	fmt.Println("Con Meo an hat")
}
func main() {
	var myCat = Cat{
		1938393,
		"Thong",
		20,
		4,
		true,
	}
	myCat.eat()
	fmt.Println(myCat)
}
