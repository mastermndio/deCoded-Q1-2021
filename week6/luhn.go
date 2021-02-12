package luhn

import (
	"strconv"
	"strings"
)

// Clean input
// disallow non digits
func validateInput(s string) bool {
	// checking if length is less than 1
	if len(s) <= 1 {
		return false
	}

	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return true
}

func Valid(str string) bool {
	// Strip all spaces
	cleanString := strings.ReplaceAll(str, " ", "")

	if validateInput(cleanString) == false {
		return false
	}

	var convertedNums []int

	for _, v := range cleanString {
		convertedNums = append(convertedNums, int(v-'0'))
	}

	for i := len(convertedNums) - 2; i >= 0; i -= 2 {
		if convertedNums[i]*2 > 9 {
			convertedNums[i] = convertedNums[i]*2 - 9
		} else {
			convertedNums[i] = convertedNums[i] * 2
		}
	}

	total := 0
	for _, v := range convertedNums {
		total += v
	}

	if total%10 == 0 {
		return true
	}

	return false
}
