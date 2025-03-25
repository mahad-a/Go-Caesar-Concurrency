package main

import (
	"fmt"
	"time"
)

/*
Passes the list of messages to a go function that encrypts each message and send each
resulting encrypted message to a channel
*/
func accelerateConcurrency() {
	// the list of messages
	messages := []string{"SYSC3101", "SYSC2100", "2 Paradigms",
		"Functional is 1st", "Concurrent is 2nd", "Winter 2025",
		"carleton.ca", "brightspace.carleton.ca/d2l/home/224508", "1125 Colonel By Dr"}

	// create the channels
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	// get the length of a third of the list
	length := len(messages) / 3
	firstPart := length        // the first third of the list
	secondPart := 2 * (length) // the second third of the list

	// set up a goroutine to send the encrypted messages to the channel
	go CaesarCipherList(messages[:firstPart], 2, channel1)           // 0/3 to 1/3
	go CaesarCipherList(messages[firstPart:secondPart], 2, channel2) // 1/3 to 2/3
	go CaesarCipherList(messages[secondPart:], 2, channel3)          // 2/3 to 3/3

	// loop through the messages and print the results
	for i := 0; i < length; i++ {
		fmt.Println(<-channel1)
		fmt.Println(<-channel2)
		fmt.Println(<-channel3)
		time.Sleep(time.Second) // synchronized
	}

}
