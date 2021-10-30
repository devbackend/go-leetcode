package medium

import (
	"strconv"
	"strings"
)

// GroupAnagrams for https://leetcode.com/problems/group-anagrams/
func GroupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		letterCounts := make([]int, 26)

		for _, c := range str {
			ix := c - 'a'
			letterCounts[ix]++
		}

		letterStrings := make([]string, 26)
		for i, letter := range letterCounts {
			letterStrings[i] = strconv.Itoa(letter)
		}

		key := strings.Join(letterStrings, "|")

		anagrams[key] = append(anagrams[key], str)
	}

	res := make([][]string, 0, len(anagrams))

	for _, items := range anagrams {
		res = append(res, items)
	}

	return res
}
