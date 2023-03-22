package util

import (
	"strconv"
)

func Atoi64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		num = 0
	}

	return num
}
