package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello")
	}
}

func main() {
	go sayHello()
	time.Sleep(500 * time.Millisecond)
}
