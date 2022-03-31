package medium

import (
	"strconv"
	"strings"
)

// CountAndSay for https://leetcode.com/problems/count-and-say/
func CountAndSay(n int) string {
	res := "1"

	for i := 1; i < n; i++ {
		prevNumSyms := strings.Split(res, "")

		var sym string

		var cnt int

		var items []string

		for k := range prevNumSyms {
			if prevNumSyms[k] == sym {
				cnt++
				continue
			}

			if sym != "" {
				items = append(items, strconv.Itoa(cnt), sym)
			}

			sym = prevNumSyms[k]
			cnt = 1
		}

		items = append(items, strconv.Itoa(cnt), sym)

		res = strings.Join(items, "")
	}

	return res
}
