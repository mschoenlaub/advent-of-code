package utils

import "sort"

func Priority(r rune) int {
	if r > 'Z' {
		return int(r - 'a' + 1)
	} else {
		return int(r - 'A' + 27)
	}
}

func dedup(s []rune) []rune {
	x := s
	removed := 0
	for len(s) > 1 {
		if s[0] == s[1] {
			for i := 1; i < len(s); i++ {
				s[i-1] = s[i]
			}
			s = s[:len(s)-1]
			removed++
		} else {
			s = s[1:]
		}
	}
	return x[:len(x)-removed]
}

func SortAndDedupe(s []rune) []rune {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return dedup(s)
}
