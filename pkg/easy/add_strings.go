package easy

import "strconv"

// AddStrings for https://leetcode.com/problems/add-strings/
func AddStrings(num1 string, num2 string) string {
	digits1 := []byte(num1)
	digits2 := []byte(num2)

	first := len(num1) - 1
	second := len(num2) - 1

	var result string

	var sum int

	for first >= 0 || second >= 0 {
		var firstDigit, secondDigit int

		if first >= 0 {
			firstDigit, _ = strconv.Atoi(string(digits1[first]))
			first--
		}

		if second >= 0 {
			secondDigit, _ = strconv.Atoi(string(digits2[second]))
			second--
		}

		sum = firstDigit + secondDigit + sum

		result = strconv.Itoa(sum%10) + result

		sum /= 10
	}

	if sum != 0 {
		result = strconv.Itoa(sum) + result
	}

	return result
}
