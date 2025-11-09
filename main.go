package main

import (
	"fmt"
	"os"
	//tgbot "github.com/go-telegram/bot"
)

func main() {
	tokenFile, errRead := os.ReadFile("token.txt")
	if errRead != nil {
		fmt.Println("Failed to open token file")
		os.Exit(1)
	}

	fmt.Println(string(tokenFile))
}
