package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type lawn struct {
	size   int
	mowers []*mower
}

func (lawn *lawn) mow() {
	for _, mower := range lawn.mowers {
		for _, instruction := range mower.instructions {
			if instruction == "F" {
				mower.moveForward(*lawn)
			} else {
				mower.turn(instruction)
			}
		}
	}
}

func (lawn *lawn) parseInput(input io.Reader) (err error) {
	scanner := bufio.NewScanner(input)
	line := 0
	var mowerPosition string
	for scanner.Scan() {
		if line == 0 {
			lawnSize := strings.Split(scanner.Text(), " ")[0]
			lawn.size, err = strconv.Atoi(lawnSize)
			if err != nil {
				return err
			}
			line++
			continue
		}

		if mowerPosition == "" {
			mowerPosition = scanner.Text()
			continue
		}
		pos := strings.Split(mowerPosition, " ")
		if len(pos) != 3 {
			return errors.New("could not parse mower position")
		}

		xPos, err := strconv.Atoi(pos[0])
		if err != nil {
			return err
		}

		yPos, err := strconv.Atoi(pos[1])
		if err != nil {
			return err
		}

		orientation := pos[2]

		newMower := mower{xPos: xPos, yPos: yPos, orientation: orientation}
		mowerPosition = ""
		for _, instruction := range strings.Split(scanner.Text(), "") {
			newMower.instructions = append(newMower.instructions, instruction)
		}
		lawn.mowers = append(lawn.mowers, &newMower)
		line++
	}

	return nil
}
