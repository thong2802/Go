package main

import (
	"fmt"
	"time"
)

//var message = make(chan string)
var intChan = make(chan int)
var done = make(chan bool)

func main() {
	//go y1()
	//
	//// nhan
	//
	////str := <- message
	//for str := range message {
	//	time.Sleep(500 * time.Millisecond)
	//	fmt.Println(str)
	go y2()
	for  {
		select {
		case payLoad := <- intChan:
			fmt.Println(payLoad)
		case <- done:
			fmt.Println("Quit")
			return

		}
	}

	}
//}
//
//func y1()  {
//	for i := 0; i < 20; i++ {
//		message <- "Hi golang" // send a string to the message channel
//	}
//	close(message)
//}

func y2()  {
	for i := 1; i < 10; i++ {
		intChan <- i
		time.Sleep(time.Second)
	}
	done <- true
}
