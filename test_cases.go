package main

var mowTestCases = []struct {
	description                                  string
	lawnWidth, lawnHeight                        int
	mower1XPos, mower1YPos                       int
	mower1Orientation                            string
	mower1Instructions                           []string
	expectedMower1EndXPos, expectedMower1EndYPos int
	expectedMower1EndOrientation                 string
	mower2XPos, mower2YPos                       int
	mower2Orientation                            string
	mower2Instructions                           []string
	expectedMower2EndXPos, expectedMower2EndYPos int
	expectedMower2EndOrientation                 string
}{
	{
		description:                  "mowers going north only",
		lawnWidth:                    5,
		lawnHeight:                   6,
		mower1XPos:                   1,
		mower1YPos:                   2,
		mower1Orientation:            "N",
		mower1Instructions:           []string{"F", "F", "F", "F", "F", "F", "F", "F"},
		expectedMower1EndXPos:        1,
		expectedMower1EndYPos:        6,
		expectedMower1EndOrientation: "N",
		mower2XPos:                   2,
		mower2YPos:                   2,
		mower2Orientation:            "N",
		mower2Instructions:           []string{"F", "F", "F", "F", "F", "F", "F", "F"},
		expectedMower2EndXPos:        2,
		expectedMower2EndYPos:        6,
		expectedMower2EndOrientation: "N",
	},
	{
		description:                  "second mower crossing the first one",
		lawnWidth:                    5,
		lawnHeight:                   5,
		mower1XPos:                   1,
		mower1YPos:                   1,
		mower1Orientation:            "E",
		mower1Instructions:           []string{"L", "F", "F", "R", "F", "F", "R"},
		expectedMower1EndXPos:        3,
		expectedMower1EndYPos:        3,
		expectedMower1EndOrientation: "S",
		mower2XPos:                   3,
		mower2YPos:                   2,
		mower2Orientation:            "N",
		mower2Instructions:           []string{"F", "F", "F", "F", "F", "R", "F", "F"},
		expectedMower2EndXPos:        5,
		expectedMower2EndYPos:        2,
		expectedMower2EndOrientation: "E",
	},
}

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
