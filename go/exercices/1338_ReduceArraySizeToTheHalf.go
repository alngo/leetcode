package exercices

import (
    "sort"
)

func MinSetSize(arr []int) int {
    half := len(arr) / 2

    count := make(map[int]int);
    for _, num := range arr {
        count[num] += 1
    }

    values := make([]int, 0, len(count))
    for _, val := range count {
        values = append(values, val)
    }

    sort.Slice(values, func(i, j int) bool {
        return values[i] > values[j]
    })

    i := 0;
    for half > 0 {
        half -= values[i];
        i++
    }

    return i;
}
