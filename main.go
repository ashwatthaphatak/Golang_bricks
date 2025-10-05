package main

import (
	"fmt"
)

type Server struct {
	users  map[string]string // map of username to user details
	userCh chan string
}

func NewServer() *Server { // constructor function to initialize the server
	return &Server{
		users:  make(map[string]string),
		userCh: make(chan string),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {
	for {
		user := <-s.userCh
		s.users[user] = user
		fmt.Printf("Added user: %s\n", user)

	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}

func main() {
	// // make a channel of strings
	// userCh := make(chan string)

	// bufferedUserCh := make(chan string, 2) // buffered channel with capacity of 2

	// user := <-userCh

	// userCh <- "Ashwattha" // deadlock occurs here when you try to send data to the channel without a receiver

	// bufferedUserCh <- "Ashwattha" // works fine, as the channel has buffer space

}

func sendMessage(msgCh chan<- string) { //send only channel
	msgCh <- "Hello, World!"
}

func readMessage(msgCh <-chan string) { //receive only channel
	msg := <-msgCh
	fmt.Println(msg)
}
