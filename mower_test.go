package main

import (
	"testing"
)

var testCases = []struct {
	orientation string
	input       string
	expected    string
}{
	{
		orientation: "N",
		input:       "L",
		expected:    "W",
	},
	{
		orientation: "S",
		input:       "R",
		expected:    "W",
	},
}

func TestTurn(t *testing.T) {
	for _, tc := range testCases {
		m := mower{
			orientation: tc.orientation,
		}
		m.turn(tc.input)
		if m.orientation != tc.expected {
			t.Fatalf("orientation : %s, turned %s, expected %s, got %s", tc.orientation, tc.input, tc.expected, m.orientation)
		}
	}
}

func TestMoveForward(t *testing.T) {
	lawn := lawn{
		width:  5,
		height: 7,
		mowers: []*mower{&mower{
			orientation: "E",
			xPos:        5,
			yPos:        2,
		},
		},
	}
	lawn.mowers[0].moveForward(lawn)
	if lawn.mowers[0].xPos != 5 {
		t.Fatal("instruction out of the lawn but mower still moved forward")
	}

	lawn.mowers[0] = &mower{
		orientation: "N",
		xPos:        5,
		yPos:        7,
	}
	lawn.mowers[0].moveForward(lawn)
	if lawn.mowers[0].yPos != 7 {
		t.Fatal("instruction out of the lawn but mower still moved forward")
	}

	lawn.mowers[0] = &mower{
		orientation: "S",
		xPos:        5,
		yPos:        0,
	}
	lawn.mowers[0].moveForward(lawn)
	if lawn.mowers[0].yPos != 0 {
		t.Fatal("instruction out of the lawn but mower still moved forward")
	}

	lawn.mowers[0] = &mower{
		orientation: "W",
		xPos:        0,
		yPos:        0,
	}
	lawn.mowers[0].moveForward(lawn)
	if lawn.mowers[0].xPos != 0 {
		t.Fatal("instruction out of the lawn but mower still moved forward")
	}

}
