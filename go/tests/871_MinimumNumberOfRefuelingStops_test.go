package tests

import (
    "leetcode/go/exercices"
    "testing"
)

type input871 struct {
    target int
    startFuel int
    stations [][]int
}

type data871 struct {
    input input871
    expected int
}

var dataProvider871 = []data871 {
    {input871{1, 1, [][]int{}}, 0},
    {input871{100, 1, [][]int{
        {10, 100},
    }}, -1},
    {input871{100, 10, [][]int{
        {10,60},
        {20,30},
        {30,30},
        {60,40},
    }}, 2},
    {input871{200, 50, [][]int{
        {25,25},
        {50,100},
        {100,100},
        {150,40},
    }}, 2},
    {input871{100, 50, [][]int{
        {25, 30},
    }}, -1},
    {input871{1000, 83, [][]int{
        {25,27},
        {36,187},
        {140,186},
        {378,6},
        {492,202},
        {517,89},
        {579,234},
        {673,86},
        {808,53},
        {954,49},
    }}, -1},
    {input871{1000, 299, [][]int{
        {13,21},
        {26,115},
        {100,47},
        {225,99},
        {299,141},
        {444,198},
        {608,190},
        {636,157},
        {647,255},
        {841,123},
    }}, 4},
}

func Test_MinRefuelStops(t *testing.T) {
    for _, data := range dataProvider871 {
        actual := exercices.MinRefuelStops(
            data.input.target, 
            data.input.startFuel,
            data.input.stations,
        );
        if actual != data.expected {
            t.Errorf("For %v, expected %d but got %d",
            data.input, data.expected, actual)
        }
    }
}
