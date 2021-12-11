package dan

import (
	"strconv"
	"strings"
)

func ParseInt64s(input string) ([]int64, error) {
	lines := strings.Split(input, "\n")
	return ConvertInt64s(lines)
}

func ConvertInt64s(inputs []string) ([]int64, error) {
	values := make([]int64, len(inputs))
	for i, input := range inputs {
		var err error
		value, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return nil, err
		}
		values[i] = value
	}
	return values, nil
}

func ParseInts(input string) ([]int, error) {
	lines := strings.Split(input, "\n")
	return ConvertInts(lines)
}

func ConvertInts(inputs []string) ([]int, error) {
	values := make([]int, len(inputs))
	for i, input := range inputs {
		var err error
		value, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			return nil, err
		}
		values[i] = int(value)
	}
	return values, nil
}
