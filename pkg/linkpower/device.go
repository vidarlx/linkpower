package linkpower

import "math"

type device interface {
	GetDistance(station Station) int
}

type Device struct {
	PosX int
	PosY int
}

func NewDevice(x int, y int) *Device {
	return &Device{
		PosX: x,
		PosY: y,
	}
}

func (d Device) GetDistance(s *Station) float64 {
	stationX, stationY := s.GetPosition()
	distX := math.Pow(float64(stationX-d.PosX), 2)
	distY := math.Pow(float64(stationY-d.PosY), 2)
	return distX + distY
}
