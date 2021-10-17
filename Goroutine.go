package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wait = sync.WaitGroup{}
func main() {
	wait.Add(2)
	go f1()

	//time.Sleep(2 *time.Second)
	go f2()
	fmt.Println(runtime.NumGoroutine())
	wait.Wait()

}

func f1()  {
	for i := 0; i < 50 ; i++ {
		fmt.Println("Toi la tho, toi an co lan: ", i)
		time.Sleep(1200 * time.Millisecond)
		fmt.Println("=============================")
	}
	wait.Done()
}
func f2()  {
	for i := 0; i < 50 ; i++ {
		fmt.Println("Toi la cao, toi an thi lan: ", i)
		time.Sleep(900 * time.Millisecond)
	}
	wait.Done()
}
