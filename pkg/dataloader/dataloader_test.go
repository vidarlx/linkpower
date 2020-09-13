package dataloader

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataLoader_LoadTestCases(t *testing.T) {
	testCases := map[string]struct {
		filename       string
		expectedOutput *TestCases
		expectedError  error
	}{
		"invalid json provided": {
			filename:      "resources/test/invalid.json",
			expectedError: errors.New(ErrInvalidData),
		},
		"json ok, invalid struct": {
			filename:      "resources/test/invalid_struct.json",
			expectedError: errors.New(ErrInvalidData),
		},
		"perfect json": {
			filename: "resources/test/ok.json",
			expectedOutput: &TestCases{
				{
					Name:   "(0,0)",
					Device: Device{X: 0, Y: 0},
					Stations: []Station{
						{X: 1, Y: 1, R: 1},
						{X: 2, Y: 2, R: 2},
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Log(name)
		dl := NewDataLoader(tc.filename)
		jsonOutput, err := dl.LoadTestCases()

		if tc.expectedError != nil {
			assert.Error(t, err)
			assert.EqualError(t, tc.expectedError, err.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, jsonOutput)
		}
	}
}
