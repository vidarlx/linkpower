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

func (w *World) FindBestLink() (*Station, error) {
	if w.UserDevice == nil {
		return nil, errors.New(ErrNoDeviceRegistered)
	}

	if len(w.Stations) == 0 {
		return nil, errors.New(ErrNoStationsAdded)
	}

	powerTable := w.measureLinkPowerByStation()
	sort.Slice(powerTable, func(i, j int) bool {
		return powerTable[i].power > powerTable[j].power
	})

	// reverse to get higher...lower order
	//for i, j := 0, len(powerTable)-1; i < j; i, j = i+1, j-1 {
	//	powerTable[i], powerTable[j] = powerTable[j], powerTable[i]
	//}

	// all stations have 0 power
	if powerTable[0].power == 0 {
		return nil, errors.New(ErrNoConnection)
	}

	bestStation, err := w.GetStationByID(powerTable[0].ID)
	if err != nil {
		return nil, err
	}

	return bestStation, nil
}

// Measure link power on all stations
func (w *World) measureLinkPowerByStation() []powerByStation {
	powerTable := []powerByStation{}
	for _, s := range w.Stations {
		deviceStationDistance := int(w.UserDevice.GetDistance(s))
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
