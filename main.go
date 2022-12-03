package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Radio Text Secret Sender\tver 1.0\n")
	fmt.Println("\nEnter:\n")

	var text string
	fmt.Scanln(&text)
	cha := make(chan string)
	go receive(text, cha)

	for {
		msg, state := <-cha
		fmt.Println(msg)
		fmt.Println("CHANNEL OPEN:", state)
		if !state {
			fmt.Println("WARNING: CHANNEL IS NO LONGER OPEN")
			break
		}
	}
}

func receive(msg string, cha chan string) {

	//	cha <- msg + "at" + time.Now().Format("04:05")
	//	msg := "Hello  "
	cha <- "\nSENT: " + msg + " at " + time.Now().Format("01:01")
	time.Sleep(1 * time.Second)

	defer close(cha)
	fmt.Println("Close Channel?")
}
