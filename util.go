package main

import "container/ring"

func newCompass() *ring.Ring {
	var compass = ring.New(4)
	for _, o := range []string{"N", "E", "S", "W"} {
		compass.Value = o
		compass = compass.Next()
	}
	return compass
}
