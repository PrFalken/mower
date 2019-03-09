package main

import (
	"reflect"
	"strings"
	"testing"
)

var errorTestCases = []string{
	`5 5
12 N
LFLFLFLFF
3 3 E
FFRFFRFRRF`,
	`X 5
1 2 N
LFLFLFLFF
3 3 E
FFRFFRFRRF`,
	`5 5
X 2 N
LFLFLFLFF
3 3 E
FFRFFRFRRF`,
	`5 5
1 X N
LFLFLFLFF
3 3 E
FFRFFRFRRF`,
	`5 5
	1 2 N
	3 3 E
	FFRFFRFRRF`,
}

func TestParseInputError(t *testing.T) {
	for _, s := range errorTestCases {
		reader := strings.NewReader(s)
		lawn := lawn{}
		err := lawn.parseInput(reader)
		if err == nil {
			t.Fatalf("parseInput for input %q should have failed but didn't.", s)
		}
	}
}

func TestParseInputHappy(t *testing.T) {
	input := `5 5
1 2 N
LFLFLFLFF
3 3 E
FFRFFRFRRF
`
	expectedLawn := lawn{
		size: 5,
		mowers: []*mower{&mower{
			orientation:  "N",
			xPos:         1,
			yPos:         2,
			instructions: []string{"L", "F", "L", "F", "L", "F", "L", "F", "F"},
		},
		},
	}
	actualLawn := lawn{}
	reader := strings.NewReader(input)
	err := actualLawn.parseInput(reader)
	if err != nil {
		t.Errorf("Parsing error : %v", err)
	}
	if !reflect.DeepEqual(expectedLawn.mowers[0], actualLawn.mowers[0]) {
		t.Errorf("lawn.parsed input: %v, want %v.", actualLawn, expectedLawn)
	}
}

func TestMow(t *testing.T) {
	lawn := lawn{
		size: 5,
		mowers: []*mower{&mower{
			orientation:  "N",
			xPos:         1,
			yPos:         2,
			instructions: []string{"L", "F", "L", "F", "L", "F", "L", "F", "F"},
		},
		},
	}
	lawn.mow()
	mower := lawn.mowers[0]
	if mower.xPos != 1 || mower.yPos != 3 || mower.orientation != "N" {
		t.Fatalf("Wrong final position/orientation, expected 1 3 N, got %v %v %s", mower.xPos, mower.yPos, mower.orientation)
	}
}
