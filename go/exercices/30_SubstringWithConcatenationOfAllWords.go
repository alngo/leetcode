package exercices

func FindSubstring(s string, words []string) []int {
    result := []int{}
    wordLength := len(words[0])
    matchLength := wordLength * len(words)
    maxLength := len(s) - matchLength + 1;

    for i := 0; i < maxLength; i++ {
        substring := s[i : i + matchLength]
        chunks := splitToChuncks(substring, wordLength)
        if containOnce(words, chunks) {
            result = append(result, i)
        }
    }
    return result;
}

func splitToChuncks(s string, length int) []string {
    var chunks []string;
    for i := 0; i < len(s); i += length {
        chunks = append(chunks, s[i : i + length])
    }
    return chunks
}

func containOnce(words []string, arr []string) bool {
    start := 0;
    for _, word := range arr {
        if pos, ok := findFrom(word, words, start); ok {
            words = swap(words, start, pos)
            start += 1;
        } else {
            break
        }
    }
    return start == len(words)
}

func findFrom(word string, words []string, start int) (int, bool) {
    for i := start; i < len(words); i++ {
        if words[i] == word {
            return i, true
        }
    }
    return -1, false
}

func swap(words []string, dest int, orig int) []string {
    tmp := words[dest]
    words[dest] = words[orig]
    words[orig] = tmp
    return words
}

