package tests

import (
	"testing"
	"leetcode/go/exercices"
)

type Data struct {
	input []int
	expected int
}

var dataProvider = []Data {
	{[]int {2, 4}, 3},
	{[]int {2, 4, 5, 10}, 7},
	{[]int {15, 13, 22, 7, 11}, 5},
}

func Test_numFactoredBinaryTrees(t *testing.T) {
	for _, data := range dataProvider {
		actual := exercices.NumFactoredBinaryTrees(data.input)
		if actual != data.expected {
			t.Errorf("For %v, espected [%d] but got [%d]", data.input, data.expected, actual)
		}
	}
}
