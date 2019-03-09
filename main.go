package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	lawn := lawn{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = lawn.parseInput(file)
	if err != nil {
		log.Fatal(err)
	}
	lawn.mow()
	for _, mower := range lawn.mowers {
		fmt.Printf("%v %v %s\n", mower.xPos, mower.yPos, mower.orientation)
	}
}
