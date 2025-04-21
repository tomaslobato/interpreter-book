package main

import (
	"os"
	"os/user"
	"fmt"
	"intbook/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to intbook lang %s! Feel free to type in commands.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
