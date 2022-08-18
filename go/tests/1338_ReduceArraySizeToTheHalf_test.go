package tests

import (
	"leetcode/go/exercices"
	"testing"
)

type data1338 struct {
    input []int
    expected int
}

var dataProvider1338 = []data1338 {
    {[]int{3,3,3,3,5,5,5,2,2,7}, 2},
    {[]int{7,7,7,7,7,7}, 1},
}

func Test_minSetSize(t *testing.T) {
    for _, data := range dataProvider1338 {
        actual := exercices.MinSetSize(data.input);
        if actual != data.expected {
            t.Errorf("For %v, expected %d but got %d",
            data.input, data.expected, actual)
        }
    }
}
