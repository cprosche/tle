package tle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTLEStringParse(t *testing.T) {

	expected := TLE{
		Name:  "ISS (ZARYA)",
		Line1: "1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9991",
		Line2: "2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250768",

		NoradId:        25544,
		Classification: "U",

		LaunchTwoDigitYear: "98",
		LaunchNumber:       "067",
		LaunchPiece:        "A",

		EpochYear: "20",
		EpochDay:  "274.51782528",

		MeanMotionFirstDerivative:  ".00000867",
		MeanMotionSecondDerivative: "00000-0",

		BStar: "22813-4",

		EphemerisType:    "0",
		ElementSetNumber: 999,
		Line1Checksum:    1,

		InclinationDegrees:     51.6441,
		RightAscensionDegrees:  93.0000,
		Eccentricity:           0.0001400,
		PerigeeArgumentDegrees: 11.0000,
		MeanAnomalyDegrees:     349.0000,
		MeanMotion:             15.49300070,
		EpochRevolutionCount:   25076,

		Line2Checksum: 8,
	}

	t.Run("Parse normal", func(t *testing.T) {
		tle := `ISS (ZARYA)
1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9991
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250768
	`

		got, err := Parse(tle)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Parse with leading and trailing whitespaces", func(t *testing.T) {
		tle := `
		
		
ISS (ZARYA)
1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9991
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250768




	`

		got, err := Parse(tle)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Parse no name", func(t *testing.T) {
		tle := `1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9991
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250768
	`

		localExpected := expected
		localExpected.Name = ""

		got, err := Parse(tle)
		assert.Nil(t, err)
		assert.Equal(t, localExpected, got)
	})
}
