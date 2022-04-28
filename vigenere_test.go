package vigenere_test

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/jamest1234/vigenere"
	"github.com/jamest1234/vigenere/keyprovider"
)

func TestKeyProvider(t *testing.T) {

	keyword := "Fantastic"
	message := "The quick brown fox jumps over the lazy dog."

	kp, err := keyprovider.New(keyword)

	if err != nil {
		t.Fail()
	}

	fmt.Printf("Testing KeyProvider with \"%s\" and \"%s\"\n", keyword, message)

	for _, char := range message {
		if keyprovider.IsAlpha(char) {
			fmt.Print(string(kp.GetChar()))
			kp.NextChar(char)
		} else {
			fmt.Print(" ")
		}
	}

	fmt.Println()

	for _, char := range message {
		if keyprovider.IsAlpha(char) {
			fmt.Print(string(unicode.ToUpper(char)))
		} else {
			fmt.Print(string(char))
		}
	}

	fmt.Println("\nCompleted")

}

func TestVigenere(t *testing.T) {
	keyword := "Fantastic"
	message := "The quick brown fox jumps over the lazy dog."

	scrambler, err := vigenere.New(keyword)
	if err != nil {
		t.Fail()
	}

	fmt.Printf("Encoding \"%s\" using \"%s\"\n", message, scrambler.GetCurrentKeyword())

	for _, char := range message {
		if keyprovider.IsAlpha(char) {
			fmt.Print(string(unicode.ToUpper(char)))
		} else {
			fmt.Print(string(char))
		}
	}

	fmt.Println()

	var encodedMessage string

	for _, char := range message {
		encodedMessage += string(scrambler.Encode(char))
	}

	fmt.Println(encodedMessage, "\nCompleted")

	scrambler.Reset()

	fmt.Printf("Decoding \"%s\" using \"%s\"\n", encodedMessage, scrambler.GetCurrentKeyword())

	for _, char := range encodedMessage {
		if keyprovider.IsAlpha(char) {
			fmt.Print(string(unicode.ToUpper(char)))
		} else {
			fmt.Print(string(char))
		}
	}

	fmt.Println()

	var decodedMessage string

	for _, char := range encodedMessage {
		decodedMessage += string(scrambler.Decode(char))
	}

	fmt.Println(decodedMessage, "\nCompleted")
}

func TestVigenereEncodeDecodeString(t *testing.T) {
	v, err := vigenere.New("Fantastic")
	if err != nil {
		t.Fail()
	}
	original := "The quick brown fox jumps over the lazy dog."
	enc := v.EncodeString(original)
	fmt.Println(enc)
	v.Reset()
	dec := v.DecodeString(enc)
	fmt.Println(dec)

	if dec != original {
		t.Fail()
	}
}
