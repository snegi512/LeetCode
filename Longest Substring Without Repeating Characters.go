package main

func lengthOfLongestSubstring(s string) int {
	str_len := len(s)
	if str_len > 1 {
		res := 1
		for i := 2; i <= str_len; i++ {
			new_map := make(map[rune]int)
			for j := 0; j < i; j++ {
				new_map[rune(s[j])]++
			}
			if len(new_map) == i {
				res = i
				continue
			}
			for j := 1; j <= str_len-i; j++ {
				new_map[rune(s[j-1])]--
				if new_map[rune(s[j-1])] == 0 {
					delete(new_map, rune(s[j-1]))
				}
				new_map[rune(s[j+i-1])]++
				if len(new_map) == i {
					res = i
					break
				}
			}
			if res != i {
				break
			}
		}
		return res
	} else {
		return str_len
	}
}
