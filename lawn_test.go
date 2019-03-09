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
		width:  5,
		height: 5,
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

var mowTestCases = []struct {
	lawnWidth, lawnHeight                      int
	mowerXPos, mowerYPos                       int
	mowerOrientation                           string
	mowerInstructions                          []string
	expectedMowerEndXPos, expectedMowerEndYPos int
	expectedMowerEndOrientation                string
}{
	{
		lawnWidth:                   5,
		lawnHeight:                  6,
		mowerXPos:                   1,
		mowerYPos:                   2,
		mowerOrientation:            "N",
		mowerInstructions:           []string{"F", "F", "F", "F", "F", "F", "F", "F"},
		expectedMowerEndXPos:        1,
		expectedMowerEndYPos:        6,
		expectedMowerEndOrientation: "N",
	},
	{
		lawnWidth:                   3,
		lawnHeight:                  5,
		mowerXPos:                   1,
		mowerYPos:                   1,
		mowerOrientation:            "E",
		mowerInstructions:           []string{"L", "F", "F", "R", "F", "F", "F", "F", "R"},
		expectedMowerEndXPos:        3,
		expectedMowerEndYPos:        3,
		expectedMowerEndOrientation: "S",
	},
}

func TestMow(t *testing.T) {
	for _, tc := range mowTestCases {
		lawn := lawn{
			width:  tc.lawnWidth,
			height: tc.lawnHeight,
			mowers: []*mower{&mower{
				orientation:  tc.mowerOrientation,
				xPos:         tc.mowerXPos,
				yPos:         tc.mowerYPos,
				instructions: tc.mowerInstructions,
			},
			},
		}
		lawn.mow()
		mower := lawn.mowers[0]
		if mower.xPos != tc.expectedMowerEndXPos ||
			mower.yPos != tc.expectedMowerEndYPos ||
			mower.orientation != tc.expectedMowerEndOrientation {
			t.Fatalf("Wrong final position/orientation, expected %v %v %s, got %v %v %s",
				tc.expectedMowerEndXPos,
				tc.expectedMowerEndYPos,
				tc.expectedMowerEndOrientation,
				mower.xPos, mower.yPos, mower.orientation)
		}
	}
}
