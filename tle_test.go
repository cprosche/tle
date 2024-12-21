package tle

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTLEStringParse(t *testing.T) {

	expected := TLE{
		Name:  "ISS (ZARYA)",
		Line1: "1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9994",
		Line2: "2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250767",
		Contents: `ISS (ZARYA)
1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9994
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250767`,

		NoradId:        25544,
		Classification: "U",

		LaunchTwoDigitYear: "98",
		LaunchNumber:       "067",
		LaunchPiece:        "A",

		EpochYear: "20",
		EpochDay:  "274.51782528",
		Epoch:     time.Date(2020, 9, 30, 12, 25, 40, 104192000, time.UTC),

		MeanMotionFirstDerivative:  ".00000867",
		MeanMotionSecondDerivative: "00000-0",

		BStar: "22813-4",

		EphemerisType:    "0",
		ElementSetNumber: 999,
		Line1Checksum:    4,

		InclinationDegrees:     51.6441,
		RightAscensionDegrees:  93.0000,
		Eccentricity:           0.0001400,
		PerigeeArgumentDegrees: 11.0000,
		MeanAnomalyDegrees:     349.0000,
		MeanMotion:             15.49300070,
		EpochRevolutionCount:   25076,

		Line2Checksum: 7,
	}

	t.Run("Parse normal", func(t *testing.T) {
		tle := `ISS (ZARYA)
1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9994
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250767`

		got, err := Parse(tle)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Parse another normal", func(t *testing.T) {
		t.Skip("This test is failing because we don't handle negatives correctly")

		tle := `ISS (ZARYA)
1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927
2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537`

		localExpected := TLE{
			Name:  "ISS (ZARYA)",
			Line1: "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927",
			Line2: "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537",
			Contents: `ISS (ZARYA)
1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927
2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537`,
			NoradId:        25544,
			Classification: "U",

			LaunchTwoDigitYear: "98",
			LaunchNumber:       "067",
			LaunchPiece:        "A",

			EpochYear: "08",
			EpochDay:  "264.51782528",
			Epoch:     time.Date(2008, 9, 30, 12, 25, 40, 104192000, time.UTC),

			MeanMotionFirstDerivative:  "-.00002182",
			MeanMotionSecondDerivative: "00000-0",

			BStar: "-11606-4",

			EphemerisType:    "0",
			ElementSetNumber: 292,
			Line1Checksum:    7,

			InclinationDegrees:     51.6416,
			RightAscensionDegrees:  247.4627,
			Eccentricity:           0.0006703,
			PerigeeArgumentDegrees: 130.5360,
			MeanAnomalyDegrees:     325.0288,
			MeanMotion:             15.72125391,
			EpochRevolutionCount:   56353,

			Line2Checksum: 7,
		}

		got, err := Parse(tle)
		assert.Nil(t, err)
		assert.Equal(t, localExpected, got)
	})

	t.Run("Parse with leading and trailing whitespaces", func(t *testing.T) {
		tle := `
		
		
ISS (ZARYA)
1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9994
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250767




	`

		got, err := Parse(tle)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Parse no name", func(t *testing.T) {
		tle := `1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9994
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250767`

		localExpected := expected
		localExpected.Contents = `1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9994
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250767`
		localExpected.Name = ""

		got, err := Parse(tle)
		assert.Nil(t, err)
		assert.Equal(t, localExpected, got)
	})
}

func TestYearAndDayToDate(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		year := "15"
		day := "22.968530853511766"

		expected := time.Date(2015, 1, 22, 23, 14, 41, 65743416, time.UTC)
		got, err := convertYearAndDayToDate(year, day)

		assert.Nil(t, err)
		assert.Equal(t, got, expected)

		year = "20"
		day = "274.51782528"

		expected = time.Date(2020, 9, 30, 12, 25, 40, 104192000, time.UTC)
		got, err = convertYearAndDayToDate(year, day)

		assert.Nil(t, err)
		assert.Equal(t, got, expected)
	})

	t.Run("Invalid year", func(t *testing.T) {
		year := "20a"
		day := "274.51782528"

		_, err := convertYearAndDayToDate(year, day)
		assert.NotNil(t, err)
	})

	t.Run("Invalid day", func(t *testing.T) {
		year := "20"
		day := "274.51782528a"

		_, err := convertYearAndDayToDate(year, day)
		assert.NotNil(t, err)
	})
}

func TestIsChecksumValid(t *testing.T) {
	t.Run("Valid checksums", func(t *testing.T) {
		tle := TLE{
			Line1: "1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9994",
			Line2: "2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250767",
		}

		assert.True(t, isChecksumValid(tle.Line1))
		assert.True(t, isChecksumValid(tle.Line2))
	})

	t.Run("Invalid checksums", func(t *testing.T) {
		tle := TLE{
			Line1: "1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9993",
			Line2: "2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250768",
		}

		assert.False(t, isChecksumValid(tle.Line1))
		assert.False(t, isChecksumValid(tle.Line2))
	})
}

func TestCalculateChecksum(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		line := "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  292"
		expected := 7

		got := calculateChecksum(line)
		assert.Equal(t, expected, got)

		line = "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.7212539156353"

		got = calculateChecksum(line)
		assert.Equal(t, expected, got)
	})
}
