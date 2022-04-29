package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jamest1234/vigenere"
)

func main() {

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("usage: vigenere <encode/decode> <keyword>")
		return
	}

	var encode = strings.HasPrefix(strings.ToLower(args[0]), "e")

	v, err := vigenere.New(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if encode {
			char = v.Encode(char)
		} else {
			char = v.Decode(char)
		}

		fmt.Print(string(char))
	}
}
