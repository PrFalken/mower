package main

import "container/ring"

type mower struct {
	orientation  string
	xPos, yPos   int
	instructions []string
}

func (mower *mower) conflictPosition(lawn *lawn, nextX, nextY int) bool {
	for _, otherMower := range lawn.mowers {
		if otherMower.xPos == nextX && otherMower.yPos == nextY {
			return true
		}
	}
	return false
}

func (mower *mower) moveForward(lawn lawn) {
	switch mower.orientation {
	case "N":
		if mower.conflictPosition(&lawn, mower.xPos, mower.yPos+1) {
			break
		}
		if mower.yPos+1 > lawn.height {
			break
		}
		mower.yPos++
	case "S":
		if mower.conflictPosition(&lawn, mower.xPos, mower.yPos-1) {
			break
		}
		if mower.yPos-1 < 0 {
			break
		}
		mower.yPos--
	case "W":
		if mower.conflictPosition(&lawn, mower.xPos-1, mower.yPos) {
			break
		}
		if mower.xPos-1 < 0 {
			break
		}
		mower.xPos--
	case "E":
		if mower.conflictPosition(&lawn, mower.xPos+1, mower.yPos) {
			break
		}
		if mower.xPos+1 > lawn.width {
			break
		}
		mower.xPos++

	}
}

func (mower *mower) turn(instruction string) {
	var compass = ring.New(4)
	for _, o := range []string{"N", "E", "S", "W"} {
		compass.Value = o
		compass = compass.Next()
	}
	for {
		if mower.orientation == compass.Value {
			break
		}
		compass = compass.Next()
	}

	if instruction == "L" {
		compass = compass.Prev()
	}
	if instruction == "R" {
		compass = compass.Next()
	}
	mower.orientation = string(compass.Value.(string))

}
