package main

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	server := NewServer()
	server.Start()

	for i := 0; i < 10; i++ {
		go func(i int) {
			server.userCh <- fmt.Sprintf("User%d", i)
			// server.addUser(fmt.Sprintf("User%d", i))
		}(i)
	}

}
