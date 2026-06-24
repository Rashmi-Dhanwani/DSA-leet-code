func maxVowels(s string, k int) int {
    count := 0

    // first window
    for i := 0; i < k; i++ {
        if isVowel(s[i]) {
            count++
        }
    }
    maxCount := count

    // slide
    for i := k; i < len(s); i++ {
        if isVowel(s[i]) {       // incoming char enters window
            count++
        }
        if isVowel(s[i-k]) {     // outgoing char leaves window
            count--
        }
        if count > maxCount {
            maxCount = count
        }
    }

    return maxCount
}

func isVowel(c byte) bool {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
        return true
    }
    return false
}