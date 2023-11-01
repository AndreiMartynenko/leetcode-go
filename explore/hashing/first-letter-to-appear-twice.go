/*
Hashing
Example 2: 2351. First Letter to Appear Twice

Given a string s, return the first character to appear twice.
It is guaranteed that the input will have a duplicate character.
*/

package main

func repeatedCharacter(s string) byte {
	seen := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			return s[i]
		}
	}

}

func main() {

}

//First approach
