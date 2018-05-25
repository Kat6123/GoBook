package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertVolume_String(t *testing.T) {
	tt := []struct {
		v           volume
		expectedStr string
	}{
		{volume{5, litr}, "5 litr"},
		{volume{5000, mll}, "5000 mll"},
		{volume{35, volumeUnit(123)}, "35 not defined"},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expectedStr, tc.v.String(), "should be equal")
	}
}

func TestAssertVolume_convert(t *testing.T) {
	// Float values?
	tt := []struct {
		testName       string
		v              volume
		unit           volumeUnit
		expectedVolume volume
	}{
		{"to millim3", volume{5, millim3}, millim3,
			volume{5, millim3}},
		{"to millilitr", volume{5e3, millim3}, mll,
			volume{5.000000000000001, mll}},
		{"to litr", volume{5e6, millim3}, litr,
			volume{5, litr}},
		{"to metr3", volume{5e9, millim3}, metr3,
			volume{5, metr3}},
		{"to undefined unit", volume{5, millim3}, volumeUnit(123),
			volume{4.065040650406504e-11, volumeUnit(123)}},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expectedVolume, tc.v.convert(tc.unit))
		})
	}
}

func TestAssertVolume_toMass(t *testing.T) {
	tt := []struct {
		testName     string
		v            volume
		unit         massUnit
		expectedMass mass
	}{
		{"to microgr", volume{5, millim3}, microgr,
			mass{5000, microgr}},
		{"to milligr", volume{5, millim3}, milligr,
			mass{5, milligr}},
		{"to gram", volume{5, millim3}, gram,
			mass{0.005, gram}},
		{"to kilogr", volume{5, litr}, kilogr,
			mass{5, kilogr}},
		{"to centner", volume{5e2, litr}, centner,
			mass{5, centner}},
		{"to ton", volume{5, metr3}, ton,
			mass{5, ton}},
		{"to undefined unit", volume{5, millim3}, massUnit(123),
			mass{4.065040650406504e-05, massUnit(123)}},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expectedMass, tc.v.toMass(tc.unit))
		})
	}
}
