/*
HASHING
Find Players With Zero or One Losses

You are given an integer array matches
where matches[i] = [winneri, loseri] indicates that
the player winneri defeated player loseri in a match.

Return a list answer of size 2 where:

answer[0] is a list of all players that have not lost any matches.
answer[1] is a list of all players that have lost exactly one match.
The values in the two lists should be returned in increasing order.

Note:

You should only consider the players
that have played at least one match.
The testcases will be generated such
that no two matches will have the same outcome.


Example 1:

Input: matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
Output: [[1,2,10],[4,5,7,8]]
Explanation:
Players 1, 2, and 10 have not lost any matches.
Players 4, 5, 7, and 8 each have lost one match.
Players 3, 6, and 9 each have lost two matches.
Thus, answer[0] = [1,2,10] and answer[1] = [4,5,7,8].
Example 2:

Input: matches = [[2,3],[1,3],[5,4],[6,4]]
Output: [[1,2,5,6],[]]
Explanation:
Players 1, 2, 5, and 6 have not lost any matches.
Players 3 and 4 each have lost two matches.
Thus, answer[0] = [1,2,5,6] and answer[1] = [].


Constraints:

1 <= matches.length <= 105
matches[i].length == 2
1 <= winneri, loseri <= 105
winneri != loseri
All matches[i] are unique.
*/

package main

// findPlayers takes a 2D array `matches` and returns a list where the first element is a list of
// players who have not lost any matches, and the second element is a list of players who have lost exactly one match.
// The results are sorted in increasing order.
func findPlayers(matches [][]int) [][]int {
	// winCounts keeps track of the number of wins for each player
	winCounts := make(map[int]int)
	// lostMatches is a set to mark players who have lost at least one match
	lostMatches := make(map[int]bool)

	// Iterate through each match
	for _, match := range matches {
		winner, loser := match[0], match[1]
		/*
			the indices 0 and 1 are used to access the first and second elements
			of an array or slice, respectively. That's why match[0] corresponds
			to the winner, and match[1] corresponds to the loser in the context of this problem.
		*/

		// Update win counts for the winner
		winCounts[winner]++

		// Mark the loser as having lost a match
		lostMatches[loser] = true
	}

	// notLost and lostOnce will store the final results
	notLost := make([]int, 0)
	lostOnce := make([]int, 0)

	// Identify players who have not lost any matches and those who have lost exactly once
	for player := range winCounts {
		if !lostMatches[player] {
			// Player has not lost any matches
			notLost = append(notLost, player)
		} else if winCounts[player] == 1 {
			// Player has lost exactly one match
			lostOnce = append(lostOnce, player)
		}
	}

	// Sort the results in increasing order
	sortIntSlice(notLost)
	sortIntSlice(lostOnce)

	// Return the final result as a 2D array
	return [][]int{notLost, lostOnce}
}

// sortIntSlice sorts an integer slice in increasing order
func sortIntSlice(slice []int) {
	// Outer loop: Iterate over each element in the slice
	for i := 0; i < len(slice)-1; i++ {
		// Inner loop: Iterate over the remaining unsorted elements in the slice
		for j := i + 1; j < len(slice); j++ {
			// Compare adjacent elements
			if slice[i] > slice[j] {
				// Swap elements if they are in the wrong order
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

/*
func main() {
	// Example input: matches represent the winners and losers of matches
	matches := [][]int{
		{1, 2},
		{3, 1},
		{2, 4},
		{4, 5},
	}

	// Call the findPlayers function and print the result
	result := findPlayers(matches)
	fmt.Println(result) // Output: [[3 4 5] [2]]
}

*/
