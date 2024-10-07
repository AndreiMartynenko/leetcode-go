package main

import "fmt"

func checkIfPangram(sentence string) bool {
	seen := make(map[rune]bool)

	for _, char := range sentence {
		if 'a' <= char && char <= 'z' {
			seen[char] = true
		}

		return len(seen) == 26
	}
}

func main() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	result := checkIfPangram(sentence)
	fmt.Println(result)
}
