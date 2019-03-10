package main

type mower struct {
	orientation  string
	xPos, yPos   int
	instructions []string
}

func (mower *mower) conflictPosition(lawn *lawn, nextX, nextY int) bool {
	if nextY < 0 || nextY > lawn.height {
		return true
	}
	if nextX < 0 || nextX > lawn.width {
		return true
	}

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
		mower.yPos++
	case "S":
		if mower.conflictPosition(&lawn, mower.xPos, mower.yPos-1) {
			break
		}
		mower.yPos--
	case "W":
		if mower.conflictPosition(&lawn, mower.xPos-1, mower.yPos) {
			break
		}
		mower.xPos--
	case "E":
		if mower.conflictPosition(&lawn, mower.xPos+1, mower.yPos) {
			break
		}
		mower.xPos++

	}
}

func (mower *mower) turn(instruction string) {
	compass := newCompass()
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
