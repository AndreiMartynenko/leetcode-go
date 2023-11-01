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
		seen[s[i]] = true
	}

	return 0 // This line should never be reached since it's guaranteed that there will be a duplicate character.
}

//Uncomment for testing
// func main() {
// 	s := "2351"
// 	result := repeatedCharacter(s)
// 	fmt.Printf("The first character to appear twice is: %c\n", result)
// }
