package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var Dog = map[string]interface{}{
		"id"   : 1,
		"name" : "Kiki",
		"age"  : 7,
	}

	// giong Json
	var Rabbit = map[string]interface{}{
		"is"   : 2,
		"name" : "Tho nau",
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
	//fmt.Println(Rabbit)

	var Forest = map[string]interface{}{
		"Member" : []interface{}{Dog, Rabbit},
	}

	// chuyen data sang String Json
	data, err := json.Marshal(Forest)

	if err == nil {
		var myJson = string(data)
		fmt.Println(myJson)
	}


	// chuyển string data thành map
	fmt.Println("===========================")
	var payLoad = `{
      	"id"   : 1,
		"name" : "Kiki",
		"age"  : 7
	}`
    var data1  = map[string]interface{}{}
	json.Unmarshal([]byte(payLoad), &data1)

	fmt.Println(data1["id"], data1["name"])


	// chuyển String json vào struct
	fmt.Println("===========================")
	var payLoad1 = `{
      	"id"   : 1,
		"name" : "Kiki",
		"age"  : 7
	}`
	var soi = MySoi{}
	json.Unmarshal([]byte(payLoad1), &soi)
	fmt.Println(soi)

}

type MySoi struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// or
type MySoi1 struct {
	ID   int
	Name string
	Age  int
}

