package linkpower

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestWorld_GetStationByID(t *testing.T) {
	uuid1 := GenerateID()
	uuid2 := GenerateID()
	uuidNonExisting := GenerateID()
	stations := []*Station{
		{ID: uuid1},
		{ID: uuid2},
	}
	w := World{Stations: stations}

	st, err := w.GetStationByID(uuid1)
	assert.NoError(t, err)
	assert.Equal(t, uuid1, st.ID)

	st, err = w.GetStationByID(uuid2)
	assert.NoError(t, err)
	assert.Equal(t, uuid2, st.ID)

	st, err = w.GetStationByID(uuidNonExisting)
	assert.EqualError(t, errors.New(ErrNoStationWithID), err.Error())
}

func TestWorld_AddDevice(t *testing.T) {
	device := NewDevice(0, 0)
	w := NewWorld()

	w.AddDevice(device)

	assert.Equal(t, w.UserDevice, device)
}

func TestWorld_AddStation(t *testing.T) {
	u := GenerateID()
	station := Station{ID: u}

	// add two different stations
	w := NewWorld()
	w.AddStation(&station)
	err := w.AddStation(NewStation(1, 2, 3))
	assert.NoError(t, err)

	st, err := w.GetStationByID(u)
	assert.NoError(t, err)
	assert.Equal(t, u, st.ID)
	assert.Equal(t, len(w.Stations), 2)

	// add one more, with the same position
	err = w.AddStation(NewStation(1, 2, 3))
	assert.Error(t, err)
}

func TestWorld_FindBestLink(t *testing.T) {
	uuids := []string{}
	for i := 0; i < 3; i++ {
		uuids = append(uuids, GenerateID())
	}

	testCases := map[string]struct {
		stations        []*Station
		userDevice      *Device
		expectedStation string
		expectedErr     error
	}{
		"no device": {
			expectedErr: errors.New(ErrNoDeviceRegistered),
		},
		"no stations": {
			userDevice:  NewDevice(1, 2),
			expectedErr: errors.New(ErrNoStationsAdded),
		},
		"two stations - both out of reach": {
			userDevice: NewDevice(0, 0),
			stations: []*Station{
				{ID: uuids[0], PosX: 1, PosY: 1, Reach: 1},
				{ID: uuids[1], PosX: 10, PosY: 10, Reach: 99},
			},
			expectedErr: errors.New(ErrNoConnection),
		},
		"only one station": {
			userDevice: NewDevice(1, 2),
			stations: []*Station{
				{ID: uuids[0], PosX: 1, PosY: 1, Reach: 5},
			},
			expectedStation: uuids[0],
		},
		"two stations - near and far from device": {
			userDevice: NewDevice(0, 0),
			stations: []*Station{
				{ID: uuids[0], PosX: 1, PosY: 1, Reach: 100},
				{ID: uuids[1], PosX: 10, PosY: 10, Reach: 1000},
			},
			expectedStation: uuids[1],
		},
		"two stations - 1st no reach": {
			userDevice: NewDevice(0, 0),
			stations: []*Station{
				{ID: uuids[0], PosX: 1, PosY: 1, Reach: 1},
				{ID: uuids[1], PosX: 10, PosY: 10, Reach: 201},
			},
			expectedStation: uuids[1],
		},
		"three stations - 3rd best power (bigger distance, but better reach)": {
			userDevice: NewDevice(0, 0),
			stations: []*Station{
				{ID: uuids[0], PosX: 1, PosY: 1, Reach: 10},
				{ID: uuids[1], PosX: 2, PosY: 2, Reach: 10},
				{ID: uuids[2], PosX: 3, PosY: 3, Reach: 30},
			},
			expectedStation: uuids[2],
		},
		"three stations - 2nd best power (3rd is nearest, but out of reach)": {
			userDevice: NewDevice(0, 0),
			stations: []*Station{
				{ID: uuids[0], PosX: 20, PosY: 20, Reach: 1000},
				{ID: uuids[1], PosX: 15, PosY: 15, Reach: 1000},
				{ID: uuids[2], PosX: 10, PosY: 15, Reach: 100},
			},
			expectedStation: uuids[1],
		},
		// examples from file
		"(0,0)": {
			userDevice: NewDevice(0, 0),
			stations: []*Station{
				{ID: uuids[0], PosX: 0, PosY: 0, Reach: 10},
				{ID: uuids[1], PosX: 20, PosY: 20, Reach: 5},
				{ID: uuids[2], PosX: 10, PosY: 0, Reach: 12},
			},
			expectedStation: uuids[0],
		},
		"(100,100)": {
			userDevice: NewDevice(100, 100),
			stations: []*Station{
				{ID: uuids[0], PosX: 0, PosY: 0, Reach: 10},
				{ID: uuids[1], PosX: 20, PosY: 20, Reach: 5},
				{ID: uuids[2], PosX: 10, PosY: 0, Reach: 12},
			},
			expectedErr: errors.New(ErrNoConnection),
		},
		"(15,10)": {
			userDevice: NewDevice(15, 10),
			stations: []*Station{
				{ID: uuids[0], PosX: 0, PosY: 0, Reach: 10},
				{ID: uuids[1], PosX: 20, PosY: 20, Reach: 5},
				{ID: uuids[2], PosX: 10, PosY: 0, Reach: 12},
			},
			expectedErr: errors.New(ErrNoConnection),
		},
		"(18,18)": {
			userDevice: NewDevice(18, 18),
			stations: []*Station{
				{ID: uuids[0], PosX: 0, PosY: 0, Reach: 10},
				{ID: uuids[1], PosX: 20, PosY: 20, Reach: 5},
				{ID: uuids[2], PosX: 10, PosY: 0, Reach: 12},
			},
			expectedErr: errors.New(ErrNoConnection),
		},
	}

	for name, tc := range testCases {
		t.Log(name)
		w := NewWorld()
		w.Stations = tc.stations
		w.UserDevice = tc.userDevice

		link, err := w.FindBestLink()
		if tc.expectedErr != nil {
			assert.Error(t, err)
			assert.EqualError(t, tc.expectedErr, err.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStation, link.ID)
		}
	}
}
