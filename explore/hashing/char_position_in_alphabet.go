package main

import (
	"fmt"
	"strings"
)

func charPositionInAlphabet(r rune) int {
	// transform the symbol to its serial number in the alphabet
	// 'a' is 97 in ASCII
	return int(r-'a') + 1
}

func hashString(s string, mod int) int {
	total := 0
	s = strings.ToLower(s) // convert to lowercase to process only with characters a-z

	for i, r := range s {
		if r >= 'a' && r <= 'z' {
			posInAlphabet := charPositionInAlphabet(r)
			total += posInAlphabet * (i + 1) // multiply by the index of the character in the string + 1
		}
	}
	return total % mod // limit the result to a modulo operation
}

func main() {
	str := "abc"
	mod := 100 // example of limited range of hash values
	hashedValue := hashString(str, mod)
	fmt.Printf("Хэш строки '%s': %d\n", str, hashedValue)
}
