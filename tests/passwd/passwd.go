package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(0)
	fmt.Print("\nConfirm password: ")
	password2, _ := terminal.ReadPassword(0)
	fmt.Println(password, password2)
}
