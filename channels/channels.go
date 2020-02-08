package main

import "fmt"

func main() {

	//creates a new channel taking a string
	message := make(chan string)

	//creates an async function to asign hello world to the channel
	go func() { message <- "hello world" }()

	//blocks until data is in the channel
	fmt.Println(<-message)

}
