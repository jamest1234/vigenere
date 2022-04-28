package keyprovider

import (
	"fmt"
	"unicode"
)

type KeyProvider struct {
	keyword []rune
	size    int
	index   int
}

func New(keyword string) (*KeyProvider, error) {

	kp := KeyProvider{}

	err := kp.Initialise(keyword)

	return &kp, err
}

func (kp *KeyProvider) Initialise(keyword string) error {
	kp.index = 0
	kp.size = len(keyword)
	kp.keyword = make([]rune, kp.size)
	for i, char := range keyword {
		if !IsAlpha(char) {
			return fmt.Errorf("Invalid character in keyword: %s", string(char))
		}
		kp.keyword[i] = unicode.ToUpper(char)
	}
	return nil
}

func (kp *KeyProvider) GetChar() rune {
	return kp.keyword[kp.index]
}

func (kp *KeyProvider) NextChar(newChar rune) {
	kp.keyword[kp.index] = unicode.ToUpper(newChar)
	kp.index++
	kp.index %= kp.size
}

func IsAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
