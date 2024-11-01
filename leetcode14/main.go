package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"a"}

	fmt.Println(longestCommonPrefix(strs))
	n := 3.14
	fmt.Printf("type is %d", int(n))
}

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
func longestCommonPrefix(strs []string) string {
	count := 0
	for _, str := range strs {
		if len(str) == 0 {
			count++
		}
	}
	if len(strs) == 0 || count != 0 {
		return strings.Join(strs, "")
	}
	res := []string{}
	for i := 0; i < len(strs)-1; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[i][j] == strs[i+1][j] && i != j {
				res = append(res, string(strs[j][i]))
			}
		}
	}
	return strings.Join(res, "")
}
