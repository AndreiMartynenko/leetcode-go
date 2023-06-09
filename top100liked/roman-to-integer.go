/*
13. Roman to Integer
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral,
just two ones added together. 12 is written as XII, which is simply X + II.
The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right.
However, the numeral for four is not IIII. Instead, the number four is written as IV.
Because the one is before the five we subtract it making four.
The same principle applies to the number nine, which is written as IX.
There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.

Example
Input: s = "III"
Output: 3
Explanation: III = 3.

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

Related Topics
Hash Table, Math, String

*/

/*in JS
var romanToInt = function(s) {
  const romanMap = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000
  };
  let result = 0;
  for (let i = 0; i < s.length; i++) {
    const currentVal = romanMap[s[i]];
    const nextVal = romanMap[s[i + 1]];
    if (nextVal && currentVal < nextVal) {
      result -= currentVal;
    } else {
      result += currentVal;
    }
  }
  return result;
}
*/

package main

func romanToInt(s string) int {
	// Create a map to store the values of each Roman numeral
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	// Initialize a variable to store the result
	var result = 0
	// Loop through each character in the input string
	for i := 0; i < len(s); i++ {
		// Get the value of the current Roman numeral
		currentVal := romanMap[string(s[i])]
		// Get the value of the next Roman numeral, if it exists
		var nextVal int
		if i+1 < len(s) {
			nextVal = romanMap[string(s[i+1])]
		}
		// If the next value is greater than the current value, subtract the current value
		if nextVal > currentVal {
			result -= currentVal
			// Otherwise, add the current value to the result
		} else {
			result += currentVal
		}
	}
	// Return the final result
	return result
}

// func main() {
// 	result := romanToInt("III")
// 	fmt.Println(result)
// }
