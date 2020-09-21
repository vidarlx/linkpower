package linkpower

import (
	"sort"

	"github.com/pkg/errors"
)

type powerByStation struct {
	ID    string
	power float64
}

type world interface {
	AddStation(s *Station)
	AddDevice(d *Device)
	GetStationByID(ID string) Station
}

type World struct {
	Stations   []*Station
	UserDevice *Device
}

func NewWorld() *World {
	return &World{}
}

func (w *World) AddStation(s *Station) error {
	if err := w.stationExists(s); err != nil {
		return err
	}
	w.Stations = append(w.Stations, s)
	return nil
}

func (w *World) AddDevice(d *Device) {
	w.UserDevice = d
}

func (w *World) GetStationByID(ID string) (*Station, error) {
	for i, s := range w.Stations {
		if s.ID == ID {
			return w.Stations[i], nil
		}
	}
	return nil, errors.New(ErrNoStationWithID)
}

func (w *World) FindBestLink() (s *Station, power float64, err error) {
	if w.UserDevice == nil {
		return nil, 0, errors.New(ErrNoDeviceRegistered)
	}

	if len(w.Stations) == 0 {
		return nil, 0, errors.New(ErrNoStationsAdded)
	}

	powerTable := w.measureLinkPowerByStation()
	sort.Slice(powerTable, func(i, j int) bool {
		return powerTable[i].power > powerTable[j].power
	})

	// all stations have 0 power
	if powerTable[0].power == 0 {
		return nil, 0, errors.New(ErrNoConnection)
	}

	bestStation, err := w.GetStationByID(powerTable[0].ID)
	if err != nil {
		return nil, 0, err
	}

	return bestStation, powerTable[0].power, nil
}

// Measure link power on all stations
func (w *World) measureLinkPowerByStation() []powerByStation {
	powerTable := []powerByStation{}
	for _, s := range w.Stations {
		deviceStationDistance := w.UserDevice.GetDistance(s)
		powerTable = append(powerTable, powerByStation{ID: s.ID, power: s.CalculatePower(deviceStationDistance)})
	}

	return powerTable
}

func (w *World) stationExists(s *Station) error {
	for _, c := range w.Stations {
		if s.PosX == c.PosX && s.PosY == c.PosY {
			return errors.New(ErrStationExists)
		}
	}
	return nil
}
