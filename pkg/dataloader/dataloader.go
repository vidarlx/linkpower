package dataloader

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
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
	absFileLocation, err := filepath.Abs(filepath.Dir(d.File))
	absFileLocation = path.Join(absFileLocation, filepath.Base(d.File))
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

	if len(*tc) == 0 {
		return nil, errors.New(ErrInvalidData)
	}

	return tc, nil
}
