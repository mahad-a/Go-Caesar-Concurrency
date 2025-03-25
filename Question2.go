package main

import (
	"fmt"
	"time"
)

// CaesarCipherList
/*
Applies the Caesar Cipher cryptographic scheme to a list of strings
*/
func CaesarCipherList(messages []string, shift int, channel chan string) {
	for _, msg := range messages { // loop through the messages
		encrypt_message := CaesarCipher(msg, shift) // encrypt using caesar cipher
		channel <- encrypt_message                  // send the message to the channel
	}
	// once all messages are sent, close the channel
	close(channel)
}

/*
Passes the list of messages to a go function that encrypts each message and send each
resulting encrypted message to a channel
*/
func concurrent() {
	// the list of messages
	messages := []string{"SYSC3101", "SYSC2100", "2 Paradigms",
		"Functional is 1st", "Concurrent is 2nd", "Winter 2025",
		"carleton.ca", "brightspace.carleton.ca/d2l/home/224508", "1125 Colonel By Dr"}

	// create the channel
	channel := make(chan string)

	// set up a goroutine to send the encrypted messages to the channel
	go CaesarCipherList(messages[:], 2, channel) // encrypt the whole list of messages

	// loop through the messages and print the results
	for i := 0; i < len(messages); i++ {
		fmt.Println(<-channel)
		time.Sleep(time.Second) // synchronized
	}

}
