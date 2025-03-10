package square

//package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

type newType int

const SidesTriangle newType = 3
const SidesSquare newType = 4
const SidesCircle newType = 0

func CalcSquare(sideLen float64, sidesNum newType) (res float64) {

	if sidesNum == SidesTriangle {
		res = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4
	} else if sidesNum == SidesSquare {
		res = math.Pow(sideLen, 2.0)
	} else if sidesNum == SidesCircle {
		res = math.Pi * math.Pow(sideLen, 2.0)
	}
	return
}
