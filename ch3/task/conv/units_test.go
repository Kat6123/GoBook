package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertMassUnit_String(t *testing.T) {
	tt := []struct {
		unit        massUnit
		expectedStr string
	}{
		{microgr, "mcgr"},
		{milligr, "mlgr"},
		{gram, "gr"},
		{kilogr, "kg"},
		{centner, "centner"},
		{ton, "ton"},
		{massUnit(123), "not defined"},
	}

	for _, tc := range tt {
		actual := tc.unit.String()
		assert.Equal(t, tc.expectedStr, actual, "should be equal")
	}
}

func TestAssertVolumeUnit_String(t *testing.T) {
	tt := []struct {
		unit        volumeUnit
		expectedStr string
	}{
		{millim3, "mm^3"},
		{mll, "mll"},
		{litr, "litr"},
		{metr3, "m^3"},
		{volumeUnit(123), "not defined"},
	}

	for _, tc := range tt {
		actual := tc.unit.String()
		assert.Equal(t, tc.expectedStr, actual, "should be equal")
	}
}
