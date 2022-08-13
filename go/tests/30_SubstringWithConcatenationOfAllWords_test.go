package tests

import (
	"leetcode/go/exercices"
	"testing"
  "reflect"
)

type Inputs struct {
    s string
    words []string
}

type Data30 struct {
    input Inputs
    expected []int
}

var dataProvider30 = []Data30 {
    {Inputs{"barfoothefoobarman", []string{"foo","bar"}}, []int{0, 9}},
    {Inputs{"wordgoodgoodgoodbestword", []string{"word","good","best","word"}}, []int{}},
    {Inputs{"wordgoodgoodgoodbestword", []string{"word","good","best","good"}}, []int{8}},
    {Inputs{"barfoofoobarthefoobarman", []string{"bar","foo","the"}}, []int{6,9,12}},
}

func Test_findSubstring(t *testing.T) {
    for _, data := range dataProvider30 {
        actual := exercices.FindSubstring(data.input.s, data.input.words);
        if !reflect.DeepEqual(actual, data.expected) {
            t.Errorf("For %s and %v, expected %v but got %v",
            data.input.s, data.input.words, data.expected, actual)
        }
    }
}
