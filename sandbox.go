package main

import (
	"fmt"
	"github.com/jkobyp/shapes"
)

func main() {
	// ***Variable assignment
	var width int = 400
	height := 600

	// ***Calling package functions
	w := shapes.InitWindow(width, height)
	fmt.Printf("hello, world")

	// ***Creating objects
	p := shapes.Point{5, 5}
	// Create a Circle called "c" with radius 10

	// ***Method calls
	c.Move(p)
	// Move the Circle to 100, 100

	// Use the documentation to add the circle to our Window w
	//

	// ***Loops
	for w.Area.Includes(c.Location()) {
		// Animate the circle's motion
	}
	// go to tour.golang.org !
}

func gt(first, second shapes.Circle) {
	// ***if statements
	if first.Radius > second.Radius {
		return true
	} else {
		return false
	}
}

//***Struct declaration
type Triangle struct {
	// Give it fields
}

//***Method declaration
func (t Triangle) Location() shapes.Point {
	// Implement me, then see if you can add me to "w"
}
