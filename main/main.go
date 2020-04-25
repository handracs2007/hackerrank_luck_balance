// https://www.hackerrank.com/challenges/luck-balance/problem

package main

import (
	"fmt"
	"sort"
)

func luckBalance(k int32, contests [][]int32) int32 {
	var unimportantContestLuckCount int32
	var importantLuckWinCount int32
	var importantLuckLoseCount int32
	var importantContestLuck = make([]int32, 0)

	for i := 0; i < len(contests); i++ {
		// count the luck for unimportant contests
		// and store the luck for important contests
		if contests[i][1] == 0 {
			unimportantContestLuckCount += contests[i][0]
		} else {
			importantContestLuck = append(importantContestLuck, contests[i][0])
		}
	}

	// Sort the import contest luck data
	sort.Slice(importantContestLuck, func(i, j int) bool {
		return importantContestLuck[i] > importantContestLuck[j]
	})

	// Now, let's count the final luck
	for i := int32(0); i < k && i < int32(len(importantContestLuck)); i++ {
		importantLuckWinCount += importantContestLuck[i]
	}

	for i := k; i < int32(len(importantContestLuck)); i++ {
		importantLuckLoseCount += importantContestLuck[i]
	}

	return importantLuckWinCount + unimportantContestLuckCount - importantLuckLoseCount
}

func main() {
	fmt.Println(luckBalance(3, [][]int32{{5, 1}, {2, 1}, {1, 1}, {8, 1}, {10, 0}, {5, 0}}))
}
