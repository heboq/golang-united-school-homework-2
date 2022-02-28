// package square
package main

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesNum int

const (
	SidesTriangle sidesNum = 3
	SidesSquare   sidesNum = 4
	SidesCircle   sidesNum = 0
)

func CalcSquare(sideLen float64, sidesNum sidesNum) (square float64) {
	squaring := math.Pow(sideLen, 2)

	switch sidesNum {
	case SidesTriangle:
		square = math.Sqrt(3) / 4 * squaring
	case SidesSquare:
		square = squaring
	case SidesCircle:
		square = math.Pi * squaring
	}
	return square
}

func main() {
	fmt.Println(CalcSquare(10.0, SidesTriangle))
	fmt.Println(CalcSquare(10.0, SidesSquare))
	fmt.Println(CalcSquare(10.0, SidesCircle))
	fmt.Println(CalcSquare(5, SidesCircle))
	fmt.Println(CalcSquare(5, 8))
}
