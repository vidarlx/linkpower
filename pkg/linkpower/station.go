package linkpower

import (
	"math"
)

type station interface {
	CalculatePower(deviceDistance int) int
}

type Station struct {
	ID    string
	PosX  int
	PosY  int
	Reach int
}

func NewStation(x int, y int, r int) *Station {
	return &Station{
		ID:    GenerateID(),
		PosX:  x,
		PosY:  y,
		Reach: r,
	}
}

func (s *Station) GetPosition() (int, int) {
	return s.PosX, s.PosY
}

func (s *Station) CalculatePower(deviceDistance int) float64 {
	if deviceDistance > s.Reach {
		return 0
	}

	return math.Pow(float64(s.Reach-deviceDistance), 2)

}
