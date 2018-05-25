package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertMass_convertTo(t *testing.T) {
	tt := []struct {
		testName      string
		m             mass
		unit          interface{}
		expectedValue interface{}
	}{
		{"mass to volume", mass{5, kilogr}, millim3,
			volume{5e6, millim3}},
		{"mass to mass", mass{5, gram}, milligr,
			mass{5e3, milligr}},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := tc.m.convertTo(tc.unit)

			if assert.NoError(t, err) {
				assert.Equal(t, tc.expectedValue, val)
			}

		})
	}
}

func TestAssertMassErr_convertTo(t *testing.T) {
	tt := []struct {
		testName       string
		m              mass
		unit           interface{}
		expectedErrMsg string
	}{
		{"mass to undefined unit", mass{5, gram}, 5,
			"cann't convert value to unit type int"},
		{"mass to nil", mass{5, gram}, nil,
			"cann't convert value to unit type <nil>"},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := tc.m.convertTo(tc.unit)

			if assert.EqualError(t, err, tc.expectedErrMsg) {
				assert.Nil(t, val)
			}
		})
	}
}

func TestAssertVolume_convertTo(t *testing.T) {
	tt := []struct {
		testName      string
		m             volume
		unit          interface{}
		expectedValue interface{}
	}{
		{"volume to volume", volume{5, litr}, litr,
			volume{5, litr}},
		{"volume to mass", volume{5, litr}, kilogr,
			mass{5, kilogr}},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := tc.m.convertTo(tc.unit)

			if assert.NoError(t, err) {
				assert.Equal(t, tc.expectedValue, val)
			}
		})
	}
}

func TestAssertVolumeErr_convertTo(t *testing.T) {
	tt := []struct {
		testName       string
		m              volume
		unit           interface{}
		expectedErrMsg string
	}{
		{"volume to undefined unit", volume{5, millim3}, 5,
			"cann't convert value to unit type int"},
		{"volume to nil", volume{5, millim3}, nil,
			"cann't convert value to unit type <nil>"},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := tc.m.convertTo(tc.unit)

			if assert.EqualError(t, err, tc.expectedErrMsg) {
				assert.Nil(t, val)
			}
		})
	}
}

func TestAssert_convert(t *testing.T) {
	tt := []struct {
		testName      string
		v             interface{}
		unit          interface{}
		expectedValue interface{}
	}{
		{"volume to volume", volume{5, litr}, litr,
			volume{5, litr}},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := convert(tc.v, tc.unit)

			if assert.NoError(t, err) {
				assert.Equal(t, tc.expectedValue, val)
			}
		})
	}
}

func TestAssertErr_convert(t *testing.T) {
	tt := []struct {
		testName       string
		v              interface{}
		unit           interface{}
		expectedErrMsg string
	}{
		{"undefined to mass", 5, kilogr,
			"cann't convert value of type int"},
		{"volume to undefined unit", volume{5, millim3}, 5,
			"cann't convert value to unit type int"},
		{"nil to nil", nil, nil,
			"cann't convert value of type <nil>"},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := convert(tc.v, tc.unit)

			if assert.EqualError(t, err, tc.expectedErrMsg){
				assert.Nil(t, val)
			}
		})
	}
}
