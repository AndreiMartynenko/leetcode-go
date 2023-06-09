package main

func reverseString(s []byte) {
	leftCharacter := 0
	rightCharacter := len(s) - 1
	for leftCharacter < rightCharacter {
		// Swap the characters at leftCharacter and rightCharacter positions
		s[leftCharacter], s[rightCharacter] = s[rightCharacter], s[leftCharacter]

		// Move the leftCharacter pointer towards the right
		leftCharacter++

		// Move the rightCharacter pointer towards the left
		rightCharacter--
	}
}

func main() {
	var testWord = "Hello World!"
	reverseString([]byte(testWord))
}
