package main

import (
	"reflect"
	"strings"
	"testing"
)

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

func TestMow(t *testing.T) {
	for _, tc := range mowTestCases {
		lawn := lawn{
			width:  tc.lawnWidth,
			height: tc.lawnHeight,
			mowers: []*mower{
				&mower{
					orientation:  tc.mower1Orientation,
					xPos:         tc.mower1XPos,
					yPos:         tc.mower1YPos,
					instructions: tc.mower1Instructions,
				},

				&mower{
					orientation:  tc.mower2Orientation,
					xPos:         tc.mower2XPos,
					yPos:         tc.mower2YPos,
					instructions: tc.mower2Instructions,
				}},
		}
		lawn.mow()
		mower1 := lawn.mowers[0]
		if mower1.xPos != tc.expectedMower1EndXPos ||
			mower1.yPos != tc.expectedMower1EndYPos ||
			mower1.orientation != tc.expectedMower1EndOrientation {
			t.Fatalf("%s : mower 1: wrong final position/orientation, expected %v %v %s, got %v %v %s",
				tc.description,
				tc.expectedMower1EndXPos,
				tc.expectedMower1EndYPos,
				tc.expectedMower1EndOrientation,
				mower1.xPos, mower1.yPos, mower1.orientation)
		}
		mower2 := lawn.mowers[1]
		if mower2.xPos != tc.expectedMower2EndXPos ||
			mower2.yPos != tc.expectedMower2EndYPos ||
			mower2.orientation != tc.expectedMower2EndOrientation {
			t.Fatalf("%s : mower 2: wrong final position/orientation, expected %v %v %s, got %v %v %s",
				tc.description,
				tc.expectedMower2EndXPos,
				tc.expectedMower2EndYPos,
				tc.expectedMower2EndOrientation,
				mower2.xPos, mower2.yPos, mower2.orientation)
		}
	}
}
