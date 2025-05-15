package main

import (
	"fmt"
	"time"
)

// Sending
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number:", num) // Process the number
		time.Sleep(time.Second * 1)            // Simulate some processing time
	}
}

// Receiving
func sum(resultChan chan int, num1, num2 int) {
	result := num1 + num2 // Sum the two numbers
	resultChan <- result  // Send the result to the channel
}

func task(doneChan chan bool) {
	defer func() { doneChan <- true }() // Notify that the task is done
	fmt.Println("Task is running, processing....")
}

// using this **emailChan <-chan string** we can receive message/data from the channel
// using this **emailSent chan<- bool** we can send message/data to the channel
func emailSender(emailChan <-chan string, emailSent chan<- bool) {
	defer func() { emailSent <- true }() // Notify that the email is sent
	for email := range emailChan {
		fmt.Println("Sending email to:", email) // Process the email
		time.Sleep(time.Second)                 // Simulate some processing time
	}
}

// Channels are used to communicate between goroutines
// They are used to send and receive data between goroutines

func main() {
	/*
		messageChan := make(chan string) // Create a channel to send messages
		messageChan <- "Hello, World!" // Send a message to the channel // blocking
		message := <-messageChan // Receive a message from the channel
		fmt.Println(message) // Print the message

	*/
	/*
		About above code:
			-------fatal error: all goroutines are asleep - deadlock!-----
		1. This error occurs because the main goroutine is waiting for a message to be sent to the channel.
		2. but there are no other goroutines to send a message to the channel.
		3. To fix this, we need to create a goroutine to send a message to the channel.

	*/
	// ----------------------------------------------------------
	// numChan := make(chan int)
	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// -----------------------------------------------------------
	// this is a blocking channel
	result := make(chan int)
	go sum(result, 5, 10) // Start the goroutine to sum the numbers
	res := <-result       // Receive the result from the channel
	fmt.Println("Sum:", res)

	// -----------------------------------------------------------
	// this is a blocking channel
	done := make(chan bool) // Create a channel to signal when the task is done
	go task(done)           // Start the goroutine
	// time.Sleep(time.Second * 2) // Wait for the goroutine to finish
	<-done // Wait for the goroutine to finish

	// ------------------------------------------------------------

	// this is a non-blocking/buffered channel

	emailChan := make(chan string, 100) // Create a buffered channel with a capacity of 100
	emailSent := make(chan bool)        // Create a channel to signal when the email is sent

	go emailSender(emailChan, emailSent) // Start the goroutine to send emails
	for i := 0; i < 10; i++ {
		emailChan <- fmt.Sprintf("%d@example.com", i) // Send email addresses to the channel
	}
	/*
		-> This is important to close the channel to signal that no more emails will be sent. Otherwise the goroutine will block waiting for more emails
			and the program will not exit
		-> Or a Deadloack will occur
	*/
	close(emailChan) // Close the email channel to signal that no more emails will be sent
	<-emailSent      // Close the channel to signal that no more emails will be sent
	fmt.Println("All emails are sent")

	// ------------------------------------------------------------

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "ping"
	}()

	// Receive from both channels
	for i := 0; i < 2; i++ {
		select {
		case num := <-chan1:
			fmt.Println("Received from chan1:", num)
		case str := <-chan2:
			fmt.Println("Received from chan2:", str)
		}
	}
}
