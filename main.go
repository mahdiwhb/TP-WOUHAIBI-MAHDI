package main

import (
	"fmt"
	"math"
	"sort"
)

func Ft_coin(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	count := 0
	value := 0
	for amount != 0 {
		for i := 0; i < len(coins); i++ {
			if value <= coins[i] && coins[i] <= amount {
				value = coins[i]
			}
		}
		if value == 0 {
			return -1
		} else if value <= amount {
			count++
			amount -= value
			value = 0
		}
	}

	return count
}

func Ft_missing(nums []int) int {
	nbmissing := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == len(nums)-count {
			count++
			i = 0
		} else {
			nbmissing = len(nums) - count
		}
	}
	return nbmissing
}

func Ft_non_overlap(intervals [][]int) int {

	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	remove_count := 0

	last_end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < last_end {
			remove_count++
		} else {
			last_end = intervals[i][1]
		}
	}

	return remove_count
}
func Ft_profit(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > res {
				res = prices[j] - prices[i]
			}
		}
	}
	return res
}

func Ft_min_window(s string, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}

	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	windowFreq := make(map[byte]int)
	left, right, formed := 0, 0, 0
	required := len(tFreq)

	minLen := math.MaxInt32
	minLeft := 0

	for right < len(s) {
		charRight := s[right]
		windowFreq[charRight]++

		if count, found := tFreq[charRight]; found && windowFreq[charRight] == count {
			formed++
		}

		for left <= right && formed == required {
			charLeft := s[left]

			if right-left+1 < minLen {
				minLen = right - left + 1
				minLeft = left
			}

			windowFreq[charLeft]--
			if count, found := tFreq[charLeft]; found && windowFreq[charLeft] < count {
				formed--
			}

			left++
		}

		right++
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[minLeft : minLeft+minLen]
}

func Ft_max_substring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := string(s[0])
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(res); j++ {
			if s[i] == res[j] {
				i = len(s) - 1
			}
		}
		if i != len(s)-1 {
			res += string(s[i])
		}
	}

	return len(res)
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // Résultat attendu: 3
	fmt.Println(Ft_coin([]int{2}, 3))        // Résultat attendu: -1
	fmt.Println(Ft_coin([]int{1}, 0))        // Résultat attendu: 0

	// Tester Ft_missing
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // Résultat attendu: ?
	fmt.Println(Ft_missing([]int{0, 1}))                      // Résultat attendu: ?
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // Résultat attendu: ?

	// Tester Ft_non_overlap
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // Résultat attendu: 1
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // Résultat attendu: 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // Résultat attendu: 2

	// Tester Ft_profit
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4})) // Résultat attendu: 5
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))    // Résultat attendu: 0

	// Tester Ft_max_substring
	fmt.Println(Ft_max_substring("abcabcbb")) // Résultat attendu: 3
	fmt.Println(Ft_max_substring("bbbbb"))    // Résultat attendu: 1

	// Tester Ft_min_window
	fmt.Println(Ft_min_window("ADOBECODEBANC", "ABC")) // Résultat attendu: "BANC"
	fmt.Println(Ft_min_window("a", "aa"))              // Résultat attendu: ""

}
