package main

import "testing"

func TestVolume_String(t *testing.T) {
	tt := []struct {
		v           volume
		expectedStr string
	}{
		{volume{5, litr}, "5 litr"},
		{volume{5000, mll}, "5000 mll"},
		{volume{35, volumeUnit(123)}, "35 not defined"},
	}

	for _, tc := range tt {
		tr := tc.v.String()
		if tr != tc.expectedStr {
			t.Errorf("string should be %v; got %v", tc.expectedStr, tr)
		}
	}
}

func TestVolume_convert(t *testing.T) {
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
			v := tc.v
			vConv := v.convert(tc.unit)
			if vConv != tc.expectedVolume {
				t.Fatalf("converted value should be %v; got %v", tc.expectedVolume, vConv)
			}
		})
	}
}

func TestVolume_toMass(t *testing.T) {
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
			m := tc.v.toMass(tc.unit)
			if m != tc.expectedMass {
				t.Fatalf("converted value of %s should be %s; got %s", tc.v, tc.expectedMass, m)
			}
		})
	}
}
