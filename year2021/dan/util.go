package dan

import (
	"sort"
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

func SortRunes(input string) string {
	bytes := []byte(input)
	runeSorter := func(l, r int) bool { return bytes[l] < bytes[r] }
	sort.Slice(bytes, runeSorter)
	return string(bytes)
}

func ByteArrays(lines []string) [][]byte {
	byteArrays := make([][]byte, len(lines))
	for x := 0; x < len(lines); x++ {
		byteArrays[x] = []byte(lines[x])
	}
	return byteArrays
}

type Int64Slice []int64

func (x Int64Slice) Len() int           { return len(x) }
func (x Int64Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x Int64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
