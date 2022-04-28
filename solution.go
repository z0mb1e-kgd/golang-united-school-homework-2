package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

type sideType int

const (
	SidesCircle   sideType = 0
	SidesTriangle sideType = 3
	SidesSquare   sideType = 4
)

func CalcSquare(sideLen float64, sidesNum sideType) float64 {
	if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2.0)
	} else if sidesNum == SidesTriangle {
		return math.Sqrt(3) * math.Pow(sideLen, 2.0) / 4
	} else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2.0)
	} else {
		return 0
	}
}
