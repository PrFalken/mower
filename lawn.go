package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type lawn struct {
	height, width int
	mowers        []*mower
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
			err := lawn.parseLawnSize(scanner.Text())
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

func (lawn *lawn) parseLawnSize(line string) (err error) {
	lawnSize := strings.Split(line, " ")
	if len(lawnSize) != 2 {
		return errors.New("could not parse lawn size line")
	}
	lawnWidth := lawnSize[0]
	lawnHeight := lawnSize[1]

	lawn.width, err = strconv.Atoi(lawnWidth)
	if err != nil {
		return err
	}
	lawn.height, err = strconv.Atoi(lawnHeight)
	if err != nil {
		return err
	}
	return nil
}
