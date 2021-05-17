package main

import (
	"fmt"
	"time"
)

func func1(done chan bool) {
	fmt.Println("Start func1")
	time.Sleep(4 * time.Second)
	fmt.Println("Finish func1")

}

func func2(done chan bool) {
	fmt.Println("Start func2")
	time.Sleep(4 * time.Second)
	fmt.Println("Finish func2")
}

func main() {
	fmt.Println("Start")
	done := make(chan bool)
	go func1(done)
	go func2(done)
	<-done
	fmt.Println("Finish")
}
