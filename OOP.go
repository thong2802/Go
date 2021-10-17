package main

import (
	"fmt"
	"reflect"
)

	type Animal interface {
		eat()
		run()
	}

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


	func (c Cat) hunt() {
		fmt.Println(c.name, "bat chuot")
	}

	func (c Cat) eat() {
		fmt.Println(c.name, "run eating hat")
	}
	func (c Cat) run()  {
		fmt.Println(c.name, "run")
	}

	func (r Rabbit) eat() {
		fmt.Println(r.name, "run eating grass")
	}
	func (c Rabbit) run()  {
		fmt.Println(c.name, "run")
	}

	func (r Rabbit) buildHome()  {
		fmt.Println(r.name, "Dao hang")
	}


	// trong struct khong ho tro viet ham ben trong nó nen ta phai viét ngoài rồi dùng con trỏ trỏ tới nó
	//func (c *Cat)eat()  {
	//	fmt.Println("Con Meo an hat")
	//}
	func main() {
		var myCat = Cat{
			1938393,
			"MEO DEN",
			20,
			4,
			true,
		}
		myCat.eat()
		myCat.hunt()
		myCat.run()
		fmt.Println(myCat)

		var myRabbit = Rabbit{
			1938393,
			"MEO DEN",
			20,
			4,
			true,
		}
		myRabbit.eat()
		myRabbit.buildHome()
		myRabbit.run()
		fmt.Println(myRabbit)


		printMSG("dcdcdcd", 2, 3,3,3,3)
		printMSG1("2","Thong", "Hi go lang")
	}
	
	// emty interface{} -> chấp nhận mọi kiểu dữ liệu || dấu ... tức là chấp nhận vô số biến đầu vào.
func printMSG(s ...interface{})  {
	fmt.Println()
	fmt.Println(s, reflect.TypeOf(s))
}

func printMSG1(s ...interface{})  {
	for i, v := range s {
		fmt.Println("i:", i, "v:" , v)
	}
}
