package dan

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day12Input = `F70
S4
E3
S4
L90
N4
R90
W3
F75
S5
L90
E1
S4
F98
N4
R90
S3
L90
W1
F39
W2
L90
E1
F99
S3
E5
F63
N4
F26
E1
R180
F58
N3
F4
E1
F45
E4
R90
E3
F76
S1
F22
R90
N1
W1
F76
W1
N5
E3
L180
S5
F87
W4
L90
F9
S2
F11
N4
L180
S3
R90
F92
L90
S1
E4
R90
W1
F1
S2
L90
F27
N3
E1
N1
E3
L180
S1
S5
R180
W5
W5
F60
S5
W5
L270
N3
R90
F65
S5
F53
W5
L90
N1
W5
L180
F87
W2
R180
S2
F77
N1
F81
L180
E5
N5
W4
L90
W4
L90
E3
N2
L90
W2
S1
F19
W1
F82
N4
R270
E5
L90
N3
R90
F81
L270
W3
R90
L270
N3
F53
E2
F84
R90
S2
F39
R180
N1
L90
F11
S2
W5
F20
W1
N4
R90
F76
E3
S5
E3
S5
W5
S2
L90
N3
E3
S5
F27
W1
L90
F65
W3
R180
F84
W2
N5
F43
L180
W3
F11
W2
R90
N1
R90
N5
W1
S4
N4
F88
N3
F87
W3
L90
F77
S5
F18
N4
F97
E5
S5
R90
F94
N5
L180
F8
N4
R90
W2
N2
L180
F4
R90
W4
S3
R90
F38
S3
E1
N5
F4
E3
R90
S4
F95
E5
F77
F32
W5
F3
R90
N1
W3
F96
L270
N2
E2
F30
S3
W2
R90
F57
R90
E1
R90
N5
E1
N4
W4
N1
W2
F47
N5
W3
L90
N4
F50
E3
R90
F27
N3
F78
N2
R90
F100
S3
F67
R90
N4
R90
N4
F88
S4
E2
S2
F31
S5
R90
W3
R180
W2
F97
F31
N1
L90
S4
F50
N3
W2
L180
F85
L180
E3
L90
F95
N4
L90
E1
S2
R180
N2
F19
N5
E5
S1
W5
R90
N1
L180
F76
S4
E5
S2
S5
E3
F53
L90
S3
E4
S1
E1
L90
F54
W1
S1
E2
N1
R90
S3
R90
F63
L90
W4
L90
F47
L90
E5
F23
W2
F97
E3
L90
N4
F54
W3
S4
W3
S2
F67
W1
S4
R90
S5
R90
W4
L180
L90
S4
F19
F42
S4
F91
R90
L180
F64
L180
W4
R90
F32
N3
F18
E2
L180
N4
E2
N1
E4
N4
F54
W5
F50
N3
L90
N5
R90
F100
E4
N1
E3
L90
F8
L90
E4
L270
F95
L90
F44
E5
R90
F79
N5
F61
S2
F71
L90
F4
N3
F25
L180
F7
W4
F96
R90
S1
R90
W1
F9
N2
W5
F1
R90
N2
F36
W4
R90
F96
W2
F26
S2
F28
E4
N1
F33
N5
F51
W2
S1
F40
N3
F67
E3
S2
R90
W1
S3
E3
L90
F75
E3
N5
E2
F52
E3
F7
N4
F4
S4
L90
S2
W5
F85
F7
L180
E1
L90
E2
S3
R180
N3
E2
R90
N5
F6
N2
L90
W1
R90
R90
F91
E2
N4
R90
S2
E3
S3
L90
W3
F61
S1
L90
W3
N2
E1
R180
E2
W5
R90
F65
N4
W3
F54
E1
N3
E5
L180
S4
N3
E5
R90
S3
R90
S4
W4
F31
S5
R90
N2
E3
F49
F47
W3
F79
R270
W2
F90
S3
F73
L180
F14
W4
F27
R90
F75
L90
N5
R90
N4
L90
N4
E2
S1
W1
S4
W5
W1
F7
W5
L180
E1
S1
F82
F36
N2
L90
E1
L90
S4
L180
N2
W3
F21
R270
F18
R180
F93
L90
W2
F4
E1
R90
E2
S3
W4
F30
E1
F69
W5
R90
E2
L180
S4
W1
N1
E3
L90
E3
R90
F69
R90
S2
L90
N4
F13
L90
E2
L90
N2
W2
N5
S4
F70
R90
F67
E4
F62
L270
F98
L90
E5
F15
E5
R90
W3
E2
F25
R180
F7
L180
W4
S3
F42
R180
R270
N1
R180
S2
F37
E2
F72
N5
W5
F61
F43
W3
R90
R270
N5
R270
E4
L90
W4
F31
F43
L180
S3
W4
R90
F20
E2
S5
L90
F75
R90
F52
W3
L90
N5
W5
N4
R90
F52
W3
F91
E1
N2
F81
R90
E2
L90
F24
E2
L180
E1
F55
E1
L90
E5
R90
F23
S3
R180
S3
F8
L180
S1
N3
F90
N5
W3
N4
L90
N3
W5
R90
E4
S4
F89
W3
N2
R90
F18
R180
W5
E4
F100
N4
F40
E3
S2
E2
F16
R90
S2
L180
F58
W1
F70
S1
R90
W3
L90
S4
F48
R90
W1
N5
E3
R90
E1
L90
F1
R90
N1
E3
F39
W3
R90
E3
L90
N5
R90
S3
W4
R180
E1
S3
F56
L90
F98
N2
W4
F67
R90
W3
S1
F33
R90
F42
L90
R90
E4
R90
E3
F74
E4
R270
F62
S5
L90
E4
F21`

type Direction int

const (
	East = iota
	South
	West
	North
)

func (d Direction) String() string {
	switch d {
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	case North:
		return "N"
	}
	return "INVALID"
}

func (d Direction) Turn(degrees int) Direction {
	val := int(d)
	val += (degrees / 90)
	for ; val < 0; val += 4 {
	}
	val %= 4
	return Direction(val)
}

func TestDirectionTurn(t *testing.T) {
	type td struct {
		start   Direction
		degrees int
		finish  Direction
	}
	tests := []td{
		{East, 0, East},
		{East, 90, South},
		{East, 180, West},
		{East, 270, North},
		{East, 360, East},
		{East, -90, North},
		{East, -180, West},
		{East, -270, South},
		{East, -360, East},
		{North, 0, North},
		{North, 90, East},
	}
	for i, test := range tests {
		finish := test.start.Turn(test.degrees)
		assert.Equal(t, test.finish, finish, "Test %d incorrect: %#v", i, test)
	}
}

type Ship struct {
	x, y      int
	direction Direction
	wx, wy    int
}

func (ship *Ship) ProcessA(s string) {
	value, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(fmt.Sprintf("Failed to get value from instruction: %s - %#v", s, err))
	}
	instruction := s[0:1]
	if instruction == "F" {
		switch ship.direction {
		case East:
			instruction = "E"
		case South:
			instruction = "S"
		case West:
			instruction = "W"
		case North:
			instruction = "N"
		}
	}
	switch instruction {
	case "N":
		ship.y += value
	case "S":
		ship.y -= value
	case "E":
		ship.x += value
	case "W":
		ship.x -= value
	case "R":
		ship.direction = ship.direction.Turn(value)
	case "L":
		ship.direction = ship.direction.Turn(0 - value)
	}
}

func (ship *Ship) ProcessB(s string) {
	value, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(fmt.Sprintf("Failed to get value from instruction: %s - %#v", s, err))
	}
	instruction := s[0:1]
	switch instruction {
	case "F":
		ship.x += ship.wx * value
		ship.y += ship.wy * value
	case "N":
		ship.wy += value
	case "S":
		ship.wy -= value
	case "E":
		ship.wx += value
	case "W":
		ship.wx -= value
	case "R":
		ship.RotateWaypoint(value)
	case "L":
		ship.RotateWaypoint(-value)
	}
}

func (ship *Ship) RotateWaypoint(degrees int) {
	switch degrees {
	case 90, -270:
		ship.wx, ship.wy = ship.wy, -ship.wx
	case 180, -180:
		ship.wx, ship.wy = -ship.wx, -ship.wy
	case 270, -90:
		ship.wx, ship.wy = -ship.wy, ship.wx
	default:
		panic(fmt.Sprintf("Invalid rotate: %d", degrees))
	}
}

func TestRotateWaypoint(t *testing.T) {
	type td struct{ sx, sy, degrees, fx, fy int }
	tests := []td{
		{1, 2, 90, 2, -1},
		{1, -2, 90, -2, -1},
		{-1, -2, -270, -2, 1},
		{-1, 2, -270, 2, 1},
		{1, 2, 180, -1, -2},
		{1, 2, -90, -2, 1},
		{1, -2, -90, 2, 1},
		{-1, -2, 270, 2, -1},
		{-1, 2, 270, -2, -1},
	}
	for i, test := range tests {
		ship := Ship{wx: test.sx, wy: test.sy}
		ship.RotateWaypoint(test.degrees)
		assert.EqualValues(t, Ship{wx: test.fx, wy: test.fy}, ship, "Test %d failed: %#v", i, test)
	}
}

func (ship *Ship) ProcessDirectionsA(s string) {
	for _, p := range strings.Split(s, "\n") {
		ship.ProcessA(p)
	}
}

func (ship *Ship) ProcessDirectionsB(s string) {
	for _, p := range strings.Split(s, "\n") {
		ship.ProcessB(p)
	}
}

func TestDay12a(t *testing.T) {
	var ship Ship
	ship.ProcessDirectionsA(day12Input)
	t.Logf("Ship after processing directions: %#v", ship)
	x, y := ship.x, ship.y
	if x < 0 {
		x = 0 - x
	}
	if y < 0 {
		y = 0 - y
	}
	t.Logf("Ship Manhattan distance: %d", x+y)
}

func TestDay12b(t *testing.T) {
	ship := Ship{
		wx: 10,
		wy: 1,
	}
	ship.ProcessDirectionsB(day12Input)
	t.Logf("Ship after processing directions: %#v", ship)
	x, y := ship.x, ship.y
	if x < 0 {
		x = 0 - x
	}
	if y < 0 {
		y = 0 - y
	}
	t.Logf("Ship Manhattan distance: %d", x+y)
}
