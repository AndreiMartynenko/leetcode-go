/*
121. Best Time to Buy and Sell Stock
You are given an array prices where prices[i]
is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day
to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction.
If you cannot achieve any profit, return 0.

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

In Go, math.Inf(1) represents positive infinity, while math.Inf(-1) represents negative infinity.
So math.Inf(1) is used to initialize minPrice to a very large positive value,
which is higher than any price in the prices slice,
and can be updated to a smaller value later on during the loop.

*/

package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	var minPrice = math.Inf(1)
	/*
		In Go, math.Inf(1) represents positive infinity, while math.Inf(-1) represents negative infinity.
		So math.Inf(1) is used to initialize minPrice to a very large positive value,
		which is higher than any price in the prices slice,
		and can be updated to a smaller value later on during the loop.
	*/
	var maxProfit = 0
	for i := 0; i < len(prices); i++ {
		if float64(prices[i]) < minPrice {
			minPrice = float64(prices[i])
		} else {
			maxProfit = int(math.Max(float64(maxProfit), float64(prices[i]-int(minPrice))))

		}
	}
	return maxProfit

}

// func maxProfit(prices []int) int {
// 	maxProfit := 0
// 	buy := math.MaxUint32
// 	for i := 0; i < len(prices); i++ {
// 		if prices[i] > buy {
// 			if prices[i]-buy > maxProfit {
// 				maxProfit = prices[i] - buy
// 			}
// 		} else {
// 			buy = prices[i]
// 		}
// 	}
// 	return maxProfit
// }

func main() {
	var prices = []int{7, 1, 5, 3, 6, 4}
	var result = maxProfit(prices)
	fmt.Println(result)
}
