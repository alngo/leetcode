package exercices

type Sequence struct {
    current int
    counter int
}

func (seq *Sequence) increment() {
    seq.counter += 1
    seq.current += 1
}

func IsPossible(nums []int) bool {

    var sequences []Sequence

    numsLoop: for _, num := range nums {
        for i := 0; i < len(sequences); i++ {
            if isConsecutive(sequences[i].current, num) {
                sequences[i].increment()
                continue numsLoop
            }
        }
        sequences = prepend(sequences, Sequence{num, 1})
    }

    for _, c := range sequences {
        if c.counter < 3 {
            return false
        }
    }
    return true
}

func prepend(slice []Sequence, element Sequence) []Sequence {
    return append([]Sequence{element}, slice...)
}

func isConsecutive(prev int, next int) bool {
    return (next - 1) == prev
}
