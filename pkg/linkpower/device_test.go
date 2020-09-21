package linkpower

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDevice_GetDistance(t *testing.T) {
	testCases := map[string]struct {
		station          *Station
		device           *Device
		expectedDistance float64
	}{
		"both in the center (0,0)": {
			station:          NewStation(0, 0, 10),
			device:           NewDevice(0, 0),
			expectedDistance: 0,
		},
		"same point (10,10)": {
			station:          NewStation(10, 10, 5),
			device:           NewDevice(10, 10),
			expectedDistance: 0,
		},
		"same quarter (north west)": {
			station:          NewStation(0, 0, 10),
			device:           NewDevice(3, 3),
			expectedDistance: 4.242640687119285,
		},
		"same quarter (south east)": {
			station:          NewStation(-3, -3, 10),
			device:           NewDevice(-5, -5),
			expectedDistance: 2.8284271247461903,
		},
		"different quarters (SE vs NW)": {
			station:          NewStation(-3, -3, 10),
			device:           NewDevice(5, 5),
			expectedDistance: 11.313708498984761,
		},
		"different quarters (NW vs SW)": {
			station:          NewStation(3, 3, 10),
			device:           NewDevice(-2, -2),
			expectedDistance: 7.0710678118654755,
		},
	}

	for name, tc := range testCases {
		t.Log(name)
		distance := tc.device.GetDistance(tc.station)
		assert.Equal(t, tc.expectedDistance, distance)
	}
}
