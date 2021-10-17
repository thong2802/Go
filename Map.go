package main

import "fmt"

func main() {
		var myMap = map[string]int{}
		fmt.Println(myMap)

		// c2
		var myMap2 = make(map[string]int)
		fmt.Println(myMap2)

		var point = make(map[string]float64)
		point["Thong"] = 8.0
		point["Dieu"] = 8.4
		point["Thai"] = 9.6
		point["Thuy"] = 5.9
		fmt.Println(point)

		// hoac
		var point1 = map[string]float64{
			"Thong" : 8.0,
			"Dieu" : 8.4,
			"Thai" : 9.6,
			"Thuy" : 5.9,
		}
		fmt.Println(point1)

		var Rabbit = map[string]interface{}{}
		Rabbit["name"] = "Tho rung"
		Rabbit["age"] = 6
		Rabbit["leg"] = 4
		Rabbit["tail"] = true
		fmt.Println(Rabbit)

		delete(Rabbit, "leg")
		fmt.Println(Rabbit)

		// kiem tra ton tai trong map
		value, ok := Rabbit["tail"]
		if ok {
			fmt.Println(value)
		} else {
			fmt.Println("Not found")
		}


		// giong Json
		var miss = map[string]interface{}{
			"is"   : 2,
			"name" : "Ngoc Trinh",
			"age"  : 21,
			"legs" :2,
			"tails": false,
			"point" :[]int{7,8,9},
			"Ba vong" : map[string]interface{}{
				"Ngực" : map[string]interface{}{
					"size" : 90,
					"weigth" : 2,
				},
				"Eo"   : 56,
				"Mông" : 86,
			},
		}
		fmt.Println(miss)


}
