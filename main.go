package main

import (
	"fmt"
	"monkE/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is monkE programming language!\n", user.Username)
	fmt.Printf("Feel free to type in Commands\n")
	fmt.Printf("Type /exit to exit from the repl\n")
	repl.Start(os.Stdin, os.Stdout)
}
