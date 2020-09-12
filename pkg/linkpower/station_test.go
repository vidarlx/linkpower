package linkpower

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStation_CalculatePower(t *testing.T) {
	testCases := map[string]struct {
		distance      int
		station       *Station
		expectedPower float64
	}{
		"distance < reach": {
			distance:      5,
			station:       NewStation(0, 0, 10),
			expectedPower: 25,
		},
		"distance == reach": {
			distance:      5,
			station:       NewStation(0, 0, 5),
			expectedPower: 0,
		},
		"distance > reach": {
			distance:      100,
			station:       NewStation(0, 0, 5),
			expectedPower: 0,
		},
		"full power": {
			distance:      0,
			station:       NewStation(0, 0, 10),
			expectedPower: 100,
		},
		"no distance, but no reach, so no power": {
			distance:      0,
			station:       NewStation(0, 0, 0),
			expectedPower: 0,
		},
	}

	for name, tc := range testCases {
		t.Log(name)
		power := tc.station.CalculatePower(tc.distance)
		assert.Equal(t, tc.expectedPower, power)
	}
}

func TestStation_GetPosition(t *testing.T) {
	testCases := map[string]struct {
		station   *Station
		expectedX int
		expectedY int
	}{
		"25, 25": {
			station:   NewStation(25, 25, 0),
			expectedX: 25,
			expectedY: 25,
		},
		"-5, 10": {
			station:   NewStation(-5, 10, 0),
			expectedX: -5,
			expectedY: 10,
		},
		"-15, -15": {
			station:   NewStation(-15, -15, 0),
			expectedX: -15,
			expectedY: -15,
		},
	}

	for name, tc := range testCases {
		t.Log(name)
		x, y := tc.station.GetPosition()
		assert.Equal(t, tc.expectedX, x)
		assert.Equal(t, tc.expectedY, y)
	}
}
