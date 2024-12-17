package tle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTLEStringParse(t *testing.T) {
	t.Run("Parse normal", func(t *testing.T) {
		tle := `ISS (ZARYA)
1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9991
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250768
	`

		expected := TLE{
			Name:  "ISS (ZARYA)",
			Line1: "1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9991",
			Line2: "2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250768",

			NoradId:        25544,
			Classification: "U",
		}
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

		expected := TLE{
			Name:  "ISS (ZARYA)",
			Line1: "1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9991",
			Line2: "2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250768",

			NoradId:        25544,
			Classification: "U",
		}
		got, err := Parse(tle)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Parse no name", func(t *testing.T) {
		tle := `1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9991
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250768
	`

		expected := TLE{
			Line1: "1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9991",
			Line2: "2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250768",

			NoradId:        25544,
			Classification: "U",
		}
		got, err := Parse(tle)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})
}
