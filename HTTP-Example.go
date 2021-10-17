package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	ID int
	Name string
	Age int
	Point []int
}

type ListStudent []Student
func main() {
	fmt.Println("My website is Running...")

	http.HandleFunc("/homepage", MyHomePage)
	http.HandleFunc("/about", MyAboutPage)
	http.HandleFunc("/api/music", MyMusicPage)
	http.HandleFunc("/api/student", MyStudentPage)
	http.HandleFunc("/api/students", MyListStudentPage)
	//http.HandleFunc("/api/music", MyMusicPage)
	log.Fatal(http.ListenAndServe(":2802", nil))

}

func MyListStudentPage(writer http.ResponseWriter, request *http.Request) {
	var students = ListStudent{
		Student{1, "Thong", 20, []int{8,9,10}},
		Student{2, "Dieu", 20, []int{8,9,10}},
		Student{3, "Thai", 9, []int{8,9,10}},
		Student{4, "Thuy", 15, []int{8,9,10}},

	}
	json.NewEncoder(writer).Encode(students)
}

func MyStudentPage(writer http.ResponseWriter, request *http.Request) {
	var student = Student{12, "Thong", 20, []int{8,9,10}}
	json.NewEncoder(writer).Encode(student)
}

func MyMusicPage(writer http.ResponseWriter, request *http.Request) {
	var data = map[string]interface{}{
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
	json.NewEncoder(writer).Encode(data)
}

func MyAboutPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>This is my about page</h1>")
}

func MyHomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "This is my home page")
}