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

		NoradIdStr:     "25544",
		NoradId:        25544,
		Classification: "U",

		InternationalDesignator: "98067A",
		LaunchTwoDigitYear:      "98",
		LaunchNumber:            "067",
		LaunchPiece:             "A",

		ElementSetEpoch: "20274.51782528",
		EpochYear:       "20",
		EpochDay:        "274.51782528",
		Epoch:           time.Date(2020, 9, 30, 12, 25, 40, 104192000, time.UTC),

		MeanMotionFirstDerivative:  0.00000867,
		MeanMotionSecondDerivative: 0.00000e-0,

		BStar: 0.22813e-4,

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

			NoradIdStr:     "25544",
			NoradId:        25544,
			Classification: "U",

			InternationalDesignator: "98067A",
			LaunchTwoDigitYear:      "98",
			LaunchNumber:            "067",
			LaunchPiece:             "A",

			ElementSetEpoch: "08264.51782528",
			EpochYear:       "08",
			EpochDay:        "264.51782528",
			Epoch:           time.Date(2008, 9, 20, 12, 25, 40, 104192000, time.UTC),

			MeanMotionFirstDerivative:  -0.00002182,
			MeanMotionSecondDerivative: 0.00000e-0,

			BStar: -0.11606e-4,

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

	t.Run("Parse alpha 5", func(t *testing.T) {
		tle := `ISS (ZARYA)
1 Z5544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2925
2 Z5544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563535`

		localExpected := TLE{
			Name:  "ISS (ZARYA)",
			Line1: "1 Z5544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2925",
			Line2: "2 Z5544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563535",
			Contents: `ISS (ZARYA)
1 Z5544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2925
2 Z5544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563535`,

			NoradIdStr:     "Z5544",
			NoradId:        335544,
			Classification: "U",

			InternationalDesignator: "98067A",
			LaunchTwoDigitYear:      "98",
			LaunchNumber:            "067",
			LaunchPiece:             "A",

			ElementSetEpoch: "08264.51782528",
			EpochYear:       "08",
			EpochDay:        "264.51782528",
			Epoch:           time.Date(2008, 9, 20, 12, 25, 40, 104192000, time.UTC),

			MeanMotionFirstDerivative:  -0.00002182,
			MeanMotionSecondDerivative: 0.00000e-0,

			BStar: -0.11606e-4,

			EphemerisType:    "0",
			ElementSetNumber: 292,
			Line1Checksum:    5,

			InclinationDegrees:     51.6416,
			RightAscensionDegrees:  247.4627,
			Eccentricity:           0.0006703,
			PerigeeArgumentDegrees: 130.5360,
			MeanAnomalyDegrees:     325.0288,
			MeanMotion:             15.72125391,
			EpochRevolutionCount:   56353,

			Line2Checksum: 5,
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

	t.Run("Many different TLEs will parse", func(t *testing.T) {
		tles := []string{
			`ISS (ZARYA)
1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9994
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250767`,
			`1 58465U 23185D   24363.11104941  .00028654  00000-0  95911-3 0  9999
2 58465  97.4051  67.6717 0007093 330.0760  30.0075 15.30897368 59670`,
			`1 33591U 09005A   24363.16300298  .00000451  00000-0  26450-3 0  9992
2 33591  99.0226  60.7854 0014697 104.8667 255.4134 14.13241242819028`,
			`1 38771U 12049A   24363.09173364  .00000381  00000-0  19401-3 0  9993
2 38771  98.6040  55.7087 0001839 103.7914 256.3468 14.21393899637091`,
			`1  9478U 76101A   24362.62107380 -.00000010  00000-0  00000+0 0  9998
2  9478   6.9318 311.1675 0107070 341.0025 254.4654  0.97590429118242`,
		}

		for _, raw := range tles {
			tle, err := Parse(raw)
			assert.Nil(t, err)
			assert.NotNil(t, tle)
		}
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

func TestParseBStar(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		bstar := "22813-4"
		expected := 0.000022813
		got, err := parseBStar(bstar)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Positive sign", func(t *testing.T) {
		bstar := "+22813-4"
		expected := 0.000022813
		got, err := parseBStar(bstar)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Positive space", func(t *testing.T) {
		bstar := " 22813-4"
		expected := 0.000022813
		got, err := parseBStar(bstar)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Negative", func(t *testing.T) {
		bstar := "-22813-4"
		expected := -0.000022813
		got, err := parseBStar(bstar)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})
}
