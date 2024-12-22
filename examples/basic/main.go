package main

import "github.com/cprosche/tle"

func main() {
	exampleTLE := `ISS (ZARYA)
1 25544U 98067A   20274.51782528  .00000867  00000-0  22813-4 0  9994
2 25544  51.6441  93.0000 0001400  11.0000 349.0000 15.49300070250767`
	t, err := tle.Parse(exampleTLE)
	if err != nil {
		panic(err)
	}

	println(t.Name)
}
