/*
20. Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example
Input: s = "()"
Output: true

Example
Input: s = "()[]{}"
Output: true

Example
Input: s = "(]"
Output: false

Related Topics
String, Stack
*/

package main

func isValid(s string) bool {
	// Initialize stack to store the closing brackets expected...
	var stack = []rune{}
	// Traverse each charater in input string...
	for idx := 0; idx < len(s); idx++ {
		// If open parentheses are present, append it to stack...
		if s[idx] == '{' {
			stack = append(stack, '}')
		} else if s[idx] == '[' {
			stack = append(stack, ']')
		} else if s[idx] == '(' {
			stack = append(stack, ')')
			// If a close bracket is found, check that it matches the last stored open bracket
		}
		if len(stack) == 0 {
			return false
		}
		lastOpenBracket := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop the last stored open bracket

		if (s[idx] == ')' && lastOpenBracket != '(') ||
			(s[idx] == '}' && lastOpenBracket != '{') ||
			(s[idx] == ']' && lastOpenBracket != '[') {
			return false // Mismatched open and close brackets
		}
	}
	return len(stack) == 0 // Return true if all brackets are matched and stack is empty
}

// Other approach
func isValid2(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false // If the string is empty or has odd length, it can't have balanced brackets
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{} // Use a slice as a stack to store open brackets

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r) // Push open brackets onto the stack
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false // No matching open bracket or mismatched brackets
		} else {
			stack = stack[:len(stack)-1] // Pop the last stored open bracket
		}
	}

	return len(stack) == 0 // Return true if all brackets are matched and stack is empty
}
