package tle

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type TLE struct {
	Name     string
	Line1    string
	Line2    string
	Contents string

	// line 1
	NoradId        int
	Classification string

	LaunchTwoDigitYear string
	LaunchNumber       string
	LaunchPiece        string

	EpochYear string
	EpochDay  string
	Epoch     time.Time

	// TODO: change these to float64
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

	Line2Checksum int
}

// TODO: support Alpha-5 format
// TODO: handle negatives

func Parse(txt string) (TLE, error) {
	result := TLE{}
	txt = strings.TrimSpace(txt)
	result.Contents = txt
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

	result.Classification = result.Line1[7:8]

	result.LaunchTwoDigitYear = result.Line1[9:11]
	result.LaunchNumber = result.Line1[11:14]
	result.LaunchPiece = strings.TrimSpace(result.Line1[14:17])

	result.EpochYear = result.Line1[18:20]
	result.EpochDay = result.Line1[20:32]
	result.Epoch, err = convertYearAndDayToDate(result.EpochYear, result.EpochDay)
	if err != nil {
		return TLE{}, err
	}

	result.MeanMotionFirstDerivative = result.Line1[34:43]
	result.MeanMotionSecondDerivative = result.Line1[45:52]

	result.BStar = result.Line1[54:61]

	result.EphemerisType = result.Line1[62:63]
	result.ElementSetNumber, err = strconv.Atoi(result.Line1[65:68])
	if err != nil {
		return TLE{}, err
	}

	result.Line1Checksum, err = strconv.Atoi(result.Line1[68:69])
	if err != nil {
		return TLE{}, err
	}

	// Parse the rest of the fields from the second line
	result.InclinationDegrees, err = strconv.ParseFloat(result.Line2[9:16], 64)
	if err != nil {
		return TLE{}, err
	}

	result.RightAscensionDegrees, err = strconv.ParseFloat(result.Line2[18:25], 64)
	if err != nil {
		return TLE{}, err
	}

	// The eccentricity is a bit tricky because it's a number with an implicit decimal point
	result.Eccentricity, err = strconv.ParseFloat("."+result.Line2[26:33], 64)
	if err != nil {
		return TLE{}, err
	}

	result.PerigeeArgumentDegrees, err = strconv.ParseFloat(result.Line2[35:42], 64)
	if err != nil {
		return TLE{}, err
	}

	result.MeanAnomalyDegrees, err = strconv.ParseFloat(result.Line2[43:51], 64)
	if err != nil {
		return TLE{}, err
	}

	result.MeanMotion, err = strconv.ParseFloat(result.Line2[52:63], 64)
	if err != nil {
		return TLE{}, err
	}

	result.EpochRevolutionCount, err = strconv.Atoi(result.Line2[63:68])
	if err != nil {
		return TLE{}, err
	}

	result.Line2Checksum, err = strconv.Atoi(result.Line2[68:69])
	if err != nil {
		return TLE{}, err
	}

	if !isChecksumValid(result.Line1) {
		return TLE{}, errors.New("line 1 checksum is invalid")
	}

	if !isChecksumValid(result.Line2) {
		return TLE{}, errors.New("line 2 checksum is invalid")
	}

	return result, nil
}

func convertYearAndDayToDate(twoDigitYear, day string) (time.Time, error) {
	// convert the two-digit year to a four-digit year
	year, err := strconv.Atoi(twoDigitYear)
	if err != nil {
		return time.Time{}, err
	}
	if year < 57 {
		year += 2000
	} else {
		year += 1900
	}

	// convert decimal day to decimal hours
	dayFloat, err := strconv.ParseFloat(day, 64)
	if err != nil {
		return time.Time{}, err
	}

	nanosecondsFloat := dayFloat * 24.0 * 60.0 * 60.0 * 1e9
	ns := time.Duration(nanosecondsFloat)

	// subtract a day the .Date adds a day
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Add(ns).Add(-time.Hour * 24), nil
}

func isChecksumValid(line string) bool {
	providedChecksum, err := strconv.Atoi(string(line[len(line)-1]))
	if err != nil {
		return false
	}

	calculatedChecksum := calculateChecksum(line[:len(line)-1])

	return providedChecksum == calculatedChecksum
}

func calculateChecksum(line string) int {
	// The checksum is the sum of all characters in the data line, modulo 10.
	// In this formula, the following non-numeric characters are assigned the indicated values:
	// Blanks, periods, letters, '+' signs: 0
	// '-' signs: 1

	result := 0
	for _, r := range line {
		switch r {
		case '0':
			result += 0
		case '1', '-':
			result += 1
		case '2':
			result += 2
		case '3':
			result += 3
		case '4':
			result += 4
		case '5':
			result += 5
		case '6':
			result += 6
		case '7':
			result += 7
		case '8':
			result += 8
		case '9':
			result += 9
		}
	}

	return result % 10
}
