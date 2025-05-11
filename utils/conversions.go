package utils

import (
	"errors"
	"strconv"
)

func StringToUint(s string) (uint, error) {
	num, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, errors.New("invalid unsigned integer string")
	}

	return uint(num), nil
}
