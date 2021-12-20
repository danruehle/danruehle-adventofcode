package dan

import (
	"fmt"
	"strings"
	"testing"
)

var day11sample = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

var day11input = `5651341452
1381541252
1878435224
6814831535
3883547383
6473548464
1885833658
3732584752
1881546128
5121717776`

type Octopuses struct {
	grid    [][]byte
	flashes int64
}

func (o *Octopuses) normalize() {
	for row := 0; row < len(o.grid); row++ {
		for col := 0; col < len(o.grid[0]); col++ {
			o.grid[row][col] -= '0'
		}
	}
}

func (o *Octopuses) step() {
	for row := 0; row < len(o.grid); row++ {
		for col := 0; col < len(o.grid[0]); col++ {
			o.grid[row][col]++
		}
	}
	for keepGoing := true; keepGoing; {
		keepGoing = false
		for row := 0; row < len(o.grid); row++ {
			for col := 0; col < len(o.grid[0]); col++ {
				current := o.grid[row][col]
				if current > 9 && current < 50 {
					keepGoing = true
					o.flash(row, col)
				}
			}
		}
	}
	o.zero()
}

func (o *Octopuses) flash(row, col int) {
	o.flashes++
	o.grid[row][col] = 50
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			o.inc(row+x, col+y)
		}
	}
}

func (o *Octopuses) inc(row, col int) {
	if row < 0 || col < 0 || row >= len(o.grid) || col >= len(o.grid[0]) {
		return
	}
	o.grid[row][col]++
	if o.grid[row][col] > 50 {
		o.grid[row][col] = 50
	}
}

func (o *Octopuses) zero() {
	for row := 0; row < len(o.grid); row++ {
		for col := 0; col < len(o.grid[0]); col++ {
			if o.grid[row][col] >= 50 {
				o.grid[row][col] = 0
			}
		}
	}
}

func (o *Octopuses) String() string {
	s := strings.Builder{}
	for row := 0; row < len(o.grid); row++ {
		for col := 0; col < len(o.grid[0]); col++ {
			s.WriteString(fmt.Sprintf("%d", o.grid[row][col]))
		}
		s.WriteRune('\n')
	}
	return s.String()
}

func (o *Octopuses) allFlashed() bool {
	for row := 0; row < len(o.grid); row++ {
		for col := 0; col < len(o.grid[0]); col++ {
			if o.grid[row][col] != 0 {
				return false
			}
		}
	}
	return true
}

func TestDay11aSample(t *testing.T) {
	octopuses := &Octopuses{grid: ByteArrays(strings.Split(day11sample, "\n"))}
	octopuses.normalize()
	for step := 0; step < 100; step++ {
		octopuses.step()
		t.Logf("step %d:\n%s", step, octopuses.String())
	}
	t.Logf("total flashes: %d", octopuses.flashes)
}

func TestDay11a(t *testing.T) {
	octopuses := &Octopuses{grid: ByteArrays(strings.Split(day11input, "\n"))}
	octopuses.normalize()
	for step := 0; step < 100; step++ {
		octopuses.step()
		t.Logf("step %d:\n%s", step, octopuses.String())
	}
	t.Logf("total flashes: %d", octopuses.flashes)
}

func TestDay11b(t *testing.T) {
	octopuses := &Octopuses{grid: ByteArrays(strings.Split(day11input, "\n"))}
	octopuses.normalize()
	step := 0
	for !octopuses.allFlashed() {
		octopuses.step()
		step++
		t.Logf("step %d:\n%s", step, octopuses.String())
	}
	t.Logf("steps to all flashes: %d", step)
}
