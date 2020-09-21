package linkpower

import (
	"math"
)

type station interface {
	CalculatePower(deviceDistance float64) float64
}

type Station struct {
	ID    string
	PosX  int
	PosY  int
	Reach float64
}

func NewStation(x int, y int, r float64) *Station {
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

func (s *Station) CalculatePower(deviceDistance float64) float64 {
	if deviceDistance > s.Reach {
		return 0
	}

	return math.Pow(float64(s.Reach-deviceDistance), 2)

}
