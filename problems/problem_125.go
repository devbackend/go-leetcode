package problems

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re := regexp.MustCompile("[a-z0-9]")
	match := re.FindAllStringSubmatch(strings.ToLower(s), -1)

	l, r := 0, len(match)-1

	for l < r {
		if match[l][0] != match[r][0] {
			return false
		}

		l++
		r--
	}

	return true
}
