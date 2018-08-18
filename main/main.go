package main

import (
	"arrrrr-scripting/repl"
	"fmt"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the arrrrr programming language!", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
