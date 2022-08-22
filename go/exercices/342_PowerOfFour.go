package exercices

func IsPowerOfFour(n int) bool {
    return n > 0 && isDivisibleBy(n, 4)
}

func isDivisibleBy(n, x int) bool {
    for n > 1 {
        if n % x != 0 {
            return false
        }
        n /= x
    }
    return true
}
