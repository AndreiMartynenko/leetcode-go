/*
70. Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps.
In how many distinct ways can you climb to the top?

Example
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Related Topics
Math, Dynemic Programming, Memoization
*/

package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]

	}
	return dp[n]

}

func main() {
	var example = climbStairs(2)
	fmt.Println(example)
}

// Other solution
var cache = make(map[int]int)

func climbStairs2(n int) int {
	if val, ok := cache[n]; ok {
		return val
	}

	if n < 3 {
		return n
	}

	cache[n] = climbStairs2(n-1) + climbStairs2(n-2)
	return cache[n]
}

// Other solution
/*
The number of ways to get to the i'th step is much like a Fibonacci number:
ways[i] = ways[i-1] + ways[i-2], where ways[i] is the number of ways
to get on the i'th step.

So, to start the calculation process we first need
to set the first two numbers in this sequence.
If we think about it a bit, there is 1 way to get t
o the step where we already stay from the beginning (step number 0),
and 1 way to get to the next step (step number 1).

The number of ways to get on the next step is:
ways[2] = ways[0] + ways[1], and so on.


*/
func climbStairs3(n int) int {
	zero, one := 1, 1
	var next int

	for i := 1; i < n; i++ {
		next = zero + one
		zero = one
		one = next
	}

	return one
}
