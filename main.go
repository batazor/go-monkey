package main

import (
	"fmt"
	"github.com/batazor/go-monkey/repl"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in command\n")
	repl.Start(os.Stdin, os.Stdout)
}
