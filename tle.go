package tle

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type TLE struct {
	Name  string
	Line1 string
	Line2 string

	// line 1
	NoradId        int
	Classification string

	LaunchTwoDigitYear string
	LaunchNumber       string
	LaunchPiece        string

	EpochYear string
	EpochDay  string
	Epoch     time.Time

	MeanMotionFirstDerivative  string
	MeanMotionSecondDerivative string

	BStar string

	EphemerisType    string
	ElementSetNumber int
	Line1Checksum    int

	// line 2
	InclinationDegrees     float64
	RightAscensionDegrees  float64
	Eccentricity           float64
	PerigeeArgumentDegrees float64
	MeanAnomalyDegrees     float64
	MeanMotion             float64
	EpochRevolutionCount   int
	Line2Checksum          int
}

func Parse(txt string) (TLE, error) {
	result := TLE{}
	txt = strings.TrimSpace(txt)
	lines := strings.Split(txt, "\n")

	switch len(lines) {
	case 2:
		result.Line1 = lines[0]
		result.Line2 = lines[1]
	case 3:
		result.Name = lines[0]
		result.Line1 = lines[1]
		result.Line2 = lines[2]
	default:
		return TLE{}, errors.New("invalid TLE, must have 2 or 3 lines:\n" + txt)
	}

	// Parse the NORAD ID from the first line
	noradId, err := strconv.Atoi(result.Line1[2:7])
	if err != nil {
		return TLE{}, err
	}
	result.NoradId = noradId

	// TODO: should we validate the classification?
	result.Classification = result.Line1[7:8]

	result.LaunchTwoDigitYear = result.Line1[9:11]
	result.LaunchNumber = result.Line1[11:14]
	result.LaunchPiece = result.Line1[14:17]

	// TODO: finish parsing the rest of the fields

	return result, nil
}
