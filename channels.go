package main

import "fmt"

func sendData(ch chan string) {
	ch <- "Hello"
	ch <- "World"
	close(ch)
}

func main() {
	ch := make(chan string)
	go sendData(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
