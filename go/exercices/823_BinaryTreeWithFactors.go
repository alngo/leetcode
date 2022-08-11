package exercices

import (
	"sort"
)

func NumFactoredBinaryTrees(arr []int) int {
	result := 0
	mod := 1_000_000_007

	if isValidInput(arr) {
		hash := make(map[int]int)
		sort.Ints(arr)

		for i, num := range arr {
			hash[num] = 1
			for j := 0; j < i; j++ {
				candidate := arr[j]
				if num % candidate == 0 {
					hash[num] += hash[candidate] * hash[num / candidate]
					hash[num] %= mod
				}
			}
		}

		for _, count := range hash {
			result += count
		}
	}

	return result % mod
}

func isValidInput(arr []int) bool {
	return len(arr) > 0 && isPositive(arr) && isUnique(arr)
}

func isPositive(arr []int) bool {
	for _, num := range arr {
		if num < 1 {
			return false
		}
	}
	return true;
}

func isUnique(arr []int) bool {
	unique := make(map[int]int)
	for _, num := range arr {
		if _, ok := unique[num]; ok {
			return false
		}
		unique[num] = 1
	}
	return true
}
