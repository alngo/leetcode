package tests

import (
	"leetcode/go/exercices"
	"testing"
)

type data342 struct {
    input int
    expected bool
}

var dataProvider342 = []data342 {
    {16, true},
    {5, false},
    {1, true},
    {-2147483648, false},
}

func Test_IsPowerOfFour(t *testing.T) {
    for _, data := range dataProvider342 {
        actual := exercices.IsPowerOfFour(data.input);
        if actual != data.expected {
            t.Errorf("For %v, expected %v but got %v",
            data.input, data.expected, actual)
        }
    }
}
