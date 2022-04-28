package vigenere

import (
	"unicode"

	"github.com/jamest1234/vigenere/keyprovider"
)

const CHARACTERS = 26

type Vigenere struct {
	mappingTable [CHARACTERS][CHARACTERS]rune
	keyword      string
	keyProvider  *keyprovider.KeyProvider
}

func IsAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func (v *Vigenere) initialiseTable() {
	for row := 0; row < CHARACTERS; row++ {
		char := 'B' + row
		for column := 0; column < CHARACTERS; column++ {
			if char > 'Z' {
				char = 'A'
			}

			v.mappingTable[row][column] = rune(char)
			char++
		}
	}
}

func New(keyword string) *Vigenere {
	v := Vigenere{
		keyword:     keyword,
		keyProvider: keyprovider.New(keyword),
	}
	v.initialiseTable()

	return &v
}

func (v *Vigenere) GetCurrentKeyword() string {
	var result string

	for range v.keyword {
		current := v.keyProvider.GetChar()
		result += string(current)
		v.keyProvider.NextChar(current)
	}

	return result
}

func (v *Vigenere) Reset() {
	v.keyProvider.Initialise(v.keyword)
}

func (v *Vigenere) Encode(char rune) rune {
	if !IsAlpha(char) {
		return char
	}

	var letterIndex rune
	var keyIndex rune

	if unicode.IsUpper(char) {
		letterIndex = char - 65
	} else {
		letterIndex = char - 32 - 65
	}

	keyIndex = v.keyProvider.GetChar() - 65

	result := v.mappingTable[letterIndex][keyIndex]
	v.keyProvider.NextChar(char)

	if unicode.IsLower(char) {
		return unicode.ToLower(result)
	}

	return result
}

func (v *Vigenere) Decode(char rune) rune {
	if !IsAlpha(char) {
		return char
	}
	upperChar := unicode.ToUpper(char)
	keyIndex := v.keyProvider.GetChar() - 65

	var result rune

	for i := 0; i < CHARACTERS; i++ {
		if v.mappingTable[keyIndex][i] == upperChar {
			result = rune(i + 65)
			break
		}
	}

	v.keyProvider.NextChar(result)

	if unicode.IsLower(char) {
		return unicode.ToLower(result)
	}

	return result
}

func (v *Vigenere) EncodeString(str string) string {

	var result string

	for _, char := range str {
		result += string(v.Encode(char))
	}

	return result
}

func (v *Vigenere) DecodeString(str string) string {

	var result string

	for _, char := range str {
		result += string(v.Decode(char))
	}

	return result
}
