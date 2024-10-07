package main

func checkIfPangram(sentence string) bool {
	seen := make(map[rune]bool)

	for _, char := range sentence {
		seen[char] = true
	}

	return len(seen) == 26
}
