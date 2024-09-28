package main

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	s = strings.ToLower(s)
	t = strings.ToLower(t)
	s = sortString(s)
	t = sortString(t)

	if s == t {
		return true
	}
	return false
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
