package dataloader

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type dataLoader interface {
	LoadTestCases() (TestCases, error)
}

type DataLoader struct {
	File string
}

func NewDataLoader(filename string) *DataLoader {
	return &DataLoader{
		File: filename,
	}
}

func (d *DataLoader) LoadTestCases() (*TestCases, error) {
	absFileLocation, err := filepath.Abs(RootDir())
	absFileLocation = path.Join(absFileLocation, d.File)
	jsonFile, err := os.Open(absFileLocation)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var tc *TestCases
	json.Unmarshal(byteValue, &tc)

	// invalid json
	if tc == nil {
		return nil, errors.New(ErrInvalidData)
	}

	// json doesn't match model struct
	if tc != nil && len(*tc) == 0 {
		return nil, errors.New(ErrInvalidData)
	}

	return tc, nil
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	return filepath.Join(filepath.Dir(b), "../..")

}
