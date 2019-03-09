package main

import (
	"fmt"
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
		size: 5,
		mowers: []*mower{&mower{
			orientation: "E",
			xPos:        5,
			yPos:        2,
		},
		},
	}
	lawn.mowers[0].moveForward(lawn)
	fmt.Println(lawn.mowers[0].xPos)
	if lawn.mowers[0].xPos != 5 {
		t.Fatal("instruction out of the lawn but mower still moved forward")
	}
}
