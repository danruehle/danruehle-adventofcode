package dan

import (
	"sort"
	"testing"
)

var day10Input = []int{
	67,
	118,
	90,
	41,
	105,
	24,
	137,
	129,
	124,
	15,
	59,
	91,
	94,
	60,
	108,
	63,
	112,
	48,
	62,
	125,
	68,
	126,
	131,
	4,
	1,
	44,
	77,
	115,
	75,
	89,
	7,
	3,
	82,
	28,
	97,
	130,
	104,
	54,
	40,
	80,
	76,
	19,
	136,
	31,
	98,
	110,
	133,
	84,
	2,
	51,
	18,
	70,
	12,
	120,
	47,
	66,
	27,
	39,
	109,
	61,
	34,
	121,
	38,
	96,
	30,
	83,
	69,
	13,
	81,
	37,
	119,
	55,
	20,
	87,
	95,
	29,
	88,
	111,
	45,
	46,
	14,
	11,
	8,
	74,
	101,
	73,
	56,
	132,
	23,
}

func TestDay10a(t *testing.T) {
	jumps := make(map[int]int)
	s := append(day10Input, 0)
	sort.Ints(s)
	s = append(s, s[len(s)-1]+3)
	for i := 1; i < len(s); i++ {
		jump := s[i] - s[i-1]
		j, ok := jumps[jump]
		if !ok {
			j = 0
		}
		j++
		jumps[jump] = j
	}
	t.Logf("Jump distribution: %#v", jumps)
	t.Logf("Product of 1 and 3 jolt differences: %d", jumps[1]*jumps[3])
}

func TestDay10b(t *testing.T) {
	multipliers := []uint64{1, 1, 1, 2, 4, 7, 12}
	s := append(day10Input, 0)
	sort.Ints(s)
	s = append(s, s[len(s)-1]+3)
	total := uint64(1)
	for i, c := 1, 1; i < len(s); i++ {
		jump := s[i] - s[i-1]
		if jump == 3 {
			m := multipliers[c]
			total *= m
			t.Logf("Found a group of %d one jolt differences, multiplied by %d to get a new total of %d", c, m, total)
			c = 0
		}
		c++
	}
	t.Logf("There are %d ways to connect adapters", total)
}

/*
1,2,3
1,3

1,2,3,4
1,2,4
1,3,4
1,4

1,2,3,4,5
1,2,3,5
1,2,4,5
1,3,4,5
1,2,5
1,3,5
1,4,5

1,2,3,4,5,6
1,2,3,4,6
1,2,4,5,6
1,3,4,5,6
1,2,3,6
1,2,4,6
1,2,5,6
1,3,4,6
1,3,5,6
1,4,5,6
1,3,6
1,4,6
*/
