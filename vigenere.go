package vigenere

import (
	"unicode"

	"github.com/jamest1234/vigenere/keyprovider"
)

const CHARACTERS = 26

type Vigenere struct {
	keyword     string
	keyProvider *keyprovider.KeyProvider
}

func letterIndexToRune(index int32) rune {
	index = (index % CHARACTERS) + 'B'
	if index > 'Z' {
		index = 'A'
	}
	return index
}

func runeToLetterIndex(char rune) int32 {
	return unicode.ToUpper(char) - 'A'

}

func New(keyword string) (*Vigenere, error) {

	kp, err := keyprovider.New(keyword)
	if err != nil {
		return nil, err
	}

	v := Vigenere{
		keyword:     keyword,
		keyProvider: kp,
	}

	return &v, nil
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
	if !keyprovider.IsAlpha(char) {
		return char
	}

	result := letterIndexToRune(runeToLetterIndex(char) + runeToLetterIndex(v.keyProvider.GetChar()))
	v.keyProvider.NextChar(char)

	if unicode.IsLower(char) {
		return unicode.ToLower(result)
	}

	return result
}

func (v *Vigenere) Decode(char rune) rune {
	if !keyprovider.IsAlpha(char) {
		return char
	}

	upperChar := unicode.ToUpper(char)
	var result rune

	for i := int32(0); i < CHARACTERS; i++ {
		if letterIndexToRune(i+runeToLetterIndex(v.keyProvider.GetChar())) == upperChar {
			result = (i % CHARACTERS) + 'A'
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
