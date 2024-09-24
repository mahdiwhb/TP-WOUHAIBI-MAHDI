package TPWOUHAIBIMAHDI

import
"math"

func Ft_min_window(s string, t string) string {
    if len(s) == 0 || len(t) == 0 {
        return ""
    }
    
    t_freq := make(map[rune]int)
    for _, char := range t {
        t_freq[char]++
    }
    
    required := len(t_freq)
    left, right, formed := 0, 0, 0
    window_counts := make(map[rune]int)
    min_length := math.MaxInt32
    min_window := ""

    for right < len(s) {
        char := rune(s[right])
        window_counts[char]++
        
        if t_freq[char] > 0 && window_counts[char] == t_freq[char] {
            formed++
        }
        
        for left <= right && formed == required {
            char = rune(s[left])
            if right-left+1 < min_length {
                min_length = right - left + 1
                min_window = s[left:right+1]
            }
            
            window_counts[char]--
            if t_freq[char] > 0 && window_counts[char] < t_freq[char] {
                formed--
            }
            left++
        }
        
        right++
    }
    
    return min_window
}

