package dan

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var day6input = `4,3,3,5,4,1,2,1,3,1,1,1,1,1,2,4,1,3,3,1,1,1,1,2,3,1,1,1,4,1,1,2,1,2,2,1,1,1,1,1,5,1,1,2,1,1,1,1,1,1,1,1,1,3,1,1,1,1,1,1,1,1,5,1,4,2,1,1,2,1,3,1,1,2,2,1,1,1,1,1,1,1,1,1,1,4,1,3,2,2,3,1,1,1,4,1,1,1,1,5,1,1,1,5,1,1,3,1,1,2,4,1,1,3,2,4,1,1,1,1,1,5,5,1,1,1,1,1,1,4,1,1,1,3,2,1,1,5,1,1,1,1,1,1,1,5,4,1,5,1,3,4,1,1,1,1,2,1,2,1,1,1,2,2,1,2,3,5,1,1,1,1,3,5,1,1,1,2,1,1,4,1,1,5,1,4,1,2,1,3,1,5,1,4,3,1,3,2,1,1,1,2,2,1,1,1,1,4,5,1,1,1,1,1,3,1,3,4,1,1,4,1,1,3,1,3,1,1,4,5,4,3,2,5,1,1,1,1,1,1,2,1,5,2,5,3,1,1,1,1,1,3,1,1,1,1,5,1,2,1,2,1,1,1,1,2,1,1,1,1,1,1,1,3,3,1,1,5,1,3,5,5,1,1,1,2,1,2,1,5,1,1,1,1,2,1,1,1,2,1`

func TestDay6a(t *testing.T) {
	fish, err := ConvertInts(strings.Split(day6input, ","))
	require.NoError(t, err, "failed to parse fish")
	for day := 1; day <= 80; day++ {
		newFish := 0
		for i, age := range fish {
			if age == 0 {
				newFish++
				fish[i] = 6
			} else {
				fish[i] = age - 1
			}
		}
		for i := 0; i < newFish; i++ {
			fish = append(fish, 8)
		}
	}
	t.Logf("Number of fish: %d", len(fish))
}

func TestDay6b(t *testing.T) {
	fish, err := ConvertInts(strings.Split(day6input, ","))
	require.NoError(t, err, "failed to parse fish")
	var fishCount [9]int64
	for _, age := range fish {
		fishCount[age]++
	}
	for day := 1; day <= 256; day++ {
		zeros := fishCount[0]
		for age := 1; age <9; age++ {
			fishCount[age-1]=fishCount[age]
		}
		fishCount[8] = zeros
		fishCount[6] += zeros
	}
	var total int64
	for _, count := range fishCount {
		total += count
	}
	t.Logf("Number of fish: %d", total)
}
