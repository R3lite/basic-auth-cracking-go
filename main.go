package main

import (
	"github.com/charmbracelet/huh";
	"fmt"
)

var (
	url string
	username string
	usernamesFile string
	password string
	passwordsFile string
)

func main() {
	huh.NewForm(
		huh.NewGroup(
		huh.NewInput().
			Title("What url do you want to attack?").
			Prompt("?").
			Value(&url),
		),
		huh.NewGroup(
		huh.NewInput().
			Title("Enter file to usernames: ").
			Prompt("?").
			Value(&usernamesFile),
		huh.NewInput().
			Title("Enter file to passwords: ").
			Prompt("?").
			Value(&passwordsFile),
		),
		).Run()
	fmt.Printf("URL: %s\nUsernames file: %s,\nPasswords file: %s",url,usernamesFile,passwordsFile)
	}
