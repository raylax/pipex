package utils

import "strconv"

func ParseInt[T int | int8 | int16 | int32 | int64](value string) T {
	if len(value) == 0 {
		return 0
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return T(parsed)
}
