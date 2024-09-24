package TPWOUHAIBIMAHDI

func Ft_max_substring(s string) int {
	charIndex := make(map[rune]int)
	maxLength, start := 0, 0

	for i, char := range s {
			if idx, found := charIndex[char]; found && idx >= start {
					start = idx + 1
			}
			charIndex[char] = i
			if i-start+1 > maxLength {
					maxLength = i - start + 1
			}
	}
	return maxLength
}



