package utils

import "strconv"

func IsNumeric(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}
