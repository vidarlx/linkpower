package dataloader

type TestCases []TestCase

type TestCase struct {
	Name     string    `json:"name"`
	Device   Device    `json:"device"`
	Stations []Station `json:"stations"`
}

type Device struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Station struct {
	X int     `json:"x"`
	Y int     `json:"y"`
	R float64 `json:"r"`
}
