package tests

import (
	"leetcode/go/exercices"
	"testing"
)

type data659 struct {
    input []int
    expected bool
}

var dataProvider659 = []data659 {
    {[]int{1,2,3,3,4,5}, true},
    {[]int{1,2,3,3,4,4,5,5}, true},
    {[]int{1,2,3,4,4,5}, false},
}

func Test_isPossible(t *testing.T) {
    for _, data := range dataProvider659 {
        actual := exercices.IsPossible(data.input);
        if actual != data.expected {
            t.Errorf("For %v, expected %v but got %v",
            data.input, data.expected, actual)
        }
    }
}
