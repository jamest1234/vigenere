package keyprovider

import "unicode"

type KeyProvider struct {
	keyword []rune
	size    int
	index   int
}

func New(keyword string) *KeyProvider {

	kp := KeyProvider{}

	kp.Initialise(keyword)

	return &kp
}

func (kp *KeyProvider) Initialise(keyword string) {
	kp.index = 0
	kp.size = len(keyword)
	kp.keyword = make([]rune, kp.size)
	for i, char := range keyword {
		kp.keyword[i] = unicode.ToUpper(char)
	}
}

func (kp *KeyProvider) GetChar() rune {
	return kp.keyword[kp.index]
}

func (kp *KeyProvider) NextChar(newChar rune) {
	kp.keyword[kp.index] = unicode.ToUpper(newChar)
	kp.index++
	kp.index %= kp.size
}
