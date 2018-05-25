package main

const (
	micro = 1e-6
	milli = 1e-3
	kilo  = 1e3
	mega  = 1e6
	giga  = 1e9
)

// If  use float 64 error doesn't accumulate as in float32, but there was written in book that float32 is better
type massUnit float64

// iota, do not export
const (
	microgr massUnit = micro
	milligr massUnit = milli
	gram    massUnit = 1
	kilogr  massUnit = kilo
	centner massUnit = 100 * kilogr
	ton     massUnit = mega
)

type volumeUnit float64

const (
	millim3 volumeUnit = 1 / giga
	mll     volumeUnit = 1 / mega
	litr    volumeUnit = 1 / kilo
	metr3   volumeUnit = 1
)

func (unit massUnit) String() string {
	switch unit {
	case microgr:
		return "mcgr"
	case milligr:
		return "mlgr"
	case gram:
		return "gr"
	case kilogr:
		return "kg"
	case centner:
		return "centner"
	case ton:
		return "ton"
	}

	return "not defined"
}

func (unit volumeUnit) String() string {
	switch unit {
	case millim3:
		return "mm^3"
	case mll:
		return "mll"
	case litr:
		return "litr"
	case metr3:
		return "m^3"
	}

	return "not defined"
}
