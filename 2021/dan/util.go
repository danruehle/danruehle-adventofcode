package dan

import (
	"strconv"
	"strings"

)

func ParseInt64s(input string) ([]int64, error) {
	lines := strings.Split(input, "\n")
	values := make([]int64, len(lines))
	for i, line := range lines {
		var err error
		values[i], err = strconv.ParseInt(line, 10, 32)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}
